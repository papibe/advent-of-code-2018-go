package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Unit struct {
	kind        string
	index       int
	units       int
	hp          int
	weak        map[string]bool
	immune      map[string]bool
	attack      int
	attack_type string
	initiative  int
	chosed      bool
}

type Damage struct {
	damage int
	unit   *Unit
}

type Attack struct {
	unit   *Unit
	victim *Unit
}

func (u *Unit) effective_power() int {
	return u.units * u.attack
}

func (u *Unit) get_attacked(unit *Unit) {
	damage := unit.effective_power()
	_, have_weakness := u.weak[unit.attack_type]
	if have_weakness {
		damage *= 2
	}
	_, is_immune := u.immune[unit.attack_type]
	if is_immune {
		damage = 0
	}

	total_hp := u.units * u.hp
	result := total_hp - damage
	if result <= 0 {
		u.units = 0
	} else {
		u.units = result / u.hp
		if (result % u.hp) != 0 {
			u.units += 1
		}
	}
	u.chosed = false
}

func parse_system(block string, name string) []Unit {
	units := []Unit{}
	lines := strings.Split(strings.Trim(block, "\n"), "\n")

	re_units_hp := regexp.MustCompile(`(\d+) units each with (\d+) hit points`)
	re_weak := regexp.MustCompile(`weak to ([^;)]*)[;)]`)
	re_immune := regexp.MustCompile(`immune to ([^;)]*)[;)]`)
	re_attack := regexp.MustCompile(`attack that does (\d+) (\w+) damage`)
	re_initiative := regexp.MustCompile(`at initiative (\d+)`)

	for i := 1; i < len(lines); i++ {
		match := re_units_hp.FindStringSubmatch(lines[i])
		n_units, _ := strconv.Atoi(match[1])
		points, _ := strconv.Atoi(match[2])

		match = re_weak.FindStringSubmatch(lines[i])
		weaks := make(map[string]bool)
		if len(match) > 0 {
			for _, weakness := range strings.Split(match[1], ",") {
				weaks[strings.Trim(weakness, " ")] = true
			}
		}

		match = re_immune.FindStringSubmatch(lines[i])
		immunes := make(map[string]bool)
		if len(match) > 0 {
			for _, immune := range strings.Split(match[1], ",") {
				immunes[strings.Trim(immune, " ")] = true
			}
		}

		match = re_attack.FindStringSubmatch(lines[i])
		attack, _ := strconv.Atoi(match[1])
		attack_type := match[2]

		match = re_initiative.FindStringSubmatch(lines[i])
		initiative, _ := strconv.Atoi(match[1])

		units = append(units, Unit{name, i, n_units, points, weaks, immunes, attack, attack_type, initiative, false})
	}

	return units
}

func parse(filename string) ([]Unit, []Unit) {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	groups := strings.Split(string(raw_data), "\n\n")
	immunes := string(groups[0])
	infections := string(groups[1])
	return parse_system(immunes, "immunes"), parse_system(infections, "infections")

}

func sort_by_effective_power(units []*Unit) {
	sort.Slice(units, func(i, j int) bool {
		iep := units[i].effective_power()
		jep := units[j].effective_power()
		if iep == jep {
			return units[i].initiative > units[j].initiative
		} else {
			return iep > jep
		}
	})
}

func sort_damages(units []Damage) {
	sort.Slice(units, func(i, j int) bool {
		idm := units[i].damage
		jdm := units[j].damage
		if idm == jdm {
			iep := units[i].unit.effective_power()
			jep := units[j].unit.effective_power()
			if iep == jep {
				return units[i].unit.initiative > units[j].unit.initiative
			} else {
				return iep > jep
			}
		} else {
			return idm > jdm
		}
	})
}

func chose_target(unit *Unit, units []*Unit) (*Unit, int, bool) {
	damages := []Damage{}
	for i := 0; i < len(units); i++ {
		enemy_unit := units[i]
		if enemy_unit.kind == unit.kind || enemy_unit.units <= 0 || enemy_unit.chosed {
			continue
		}
		damage := unit.effective_power()
		_, have_weakness := enemy_unit.weak[unit.attack_type]
		if have_weakness {
			damage *= 2
		}
		_, is_immune := enemy_unit.immune[unit.attack_type]
		if is_immune {
			damage = 0
		}
		damages = append(damages, Damage{damage, enemy_unit})
	}

	sort_damages(damages)

	if len(damages) > 0 && damages[0].damage > 0 && (damages[0].damage/damages[0].unit.hp) > 0 {
		damages[0].unit.chosed = true
		return damages[0].unit, damages[0].damage, true
	}
	return nil, 0, false
}

func sort_units_by_initiative(attack []Attack) {
	sort.Slice(attack, func(i, j int) bool {
		return attack[i].unit.initiative > attack[j].unit.initiative
	})
}

func end_condition(immunes, infections []Unit) bool {
	all_dead := true
	for _, unit := range immunes {
		if unit.units > 0 {
			all_dead = false
			break
		}
	}

	if all_dead {
		return true
	}

	for _, unit := range infections {
		if unit.units > 0 {
			return false
		}
	}

	return true
}
func solve(immunes, infections []Unit, boost int) (int, bool) {
	all_units := []*Unit{}
	for i := 0; i < len(immunes); i++ {
		immunes[i].attack += boost
		all_units = append(all_units, &immunes[i])
	}
	for i := 0; i < len(infections); i++ {
		all_units = append(all_units, &infections[i])
	}
	for !end_condition(immunes, infections) {

		sort_by_effective_power(all_units)
		attacks := []Attack{}

		for i := 0; i < len(all_units); i++ {
			unit := all_units[i]
			victim, _, is_there_a_victim := chose_target(unit, all_units)
			if is_there_a_victim {
				attacks = append(attacks, Attack{unit, victim})
			}
		}

		//
		// Attack phase
		//
		sort_units_by_initiative(attacks)
		if len(attacks) == 0 {
			break
		}
		for i := 0; i < len(attacks); i++ {
			attack := attacks[i]
			attacks[i].victim.get_attacked(attack.unit)
		}
	}
	immunes_win := false
	units_reminding := 0
	for i := 0; i < len(immunes); i++ {
		if immunes[i].units > 0 {
			units_reminding += immunes[i].units
			immunes_win = true

		}
	}
	for i := 0; i < len(infections); i++ {
		if infections[i].units > 0 {
			units_reminding += infections[i].units
			immunes_win = false
		}
	}
	return units_reminding, immunes_win
}

func solution(filename string) int {
	immunes, infections := parse(filename)
	working_immunes := make([]Unit, len(immunes))
	working_infections := make([]Unit, len(infections))

	boost := 1
	prev_boost := 1
	for {
		copy(working_immunes, immunes)
		copy(working_infections, infections)

		_, immunes_win := solve(working_immunes, working_infections, boost)
		if immunes_win {
			break
		}
		prev_boost = boost
		boost *= 2
	}
	min_boost := prev_boost
	max_boost := boost
	var mid_boost int

	for min_boost+1 < max_boost {
		mid_boost = (min_boost + max_boost) / 2

		copy(working_immunes, immunes)
		copy(working_infections, infections)

		_, immunes_win := solve(working_immunes, working_infections, mid_boost)
		if immunes_win {
			max_boost = mid_boost
		} else {
			min_boost = mid_boost
		}
	}
	copy(working_immunes, immunes)
	copy(working_infections, infections)

	units, _ := solve(working_immunes, working_infections, max_boost)
	return units
}

func main() {
	fmt.Println(solution("./example.txt")) // 51
	fmt.Println(solution("./input.txt"))   // 5549
}
