package dndml

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	rollFmt       = "%dd%d"
	rollKeepFmt   = "%s%d"
	rollRerollFmt = "%s%d"
	rollBonusFmt  = "+%d"
)

var (
	rollRegex   = regexp.MustCompile(`^(?P<dice>\d+)d(?P<die>\d{1,3})((?P<keepScheme>(kh|kl))(?P<keepCount>[0-9]+))?(?P<reroll>((rr|ro)\d{1,3})*)(\+(?P<bonus>\d+))?$`)
	rerollRegex = regexp.MustCompile(`(?P<rerollScheme>(ro|rr))(?P<rerollDie>\d{1,3})*`)
)

func (x *Roll) UnmarshalText(text []byte) error {
	matches := rollRegex.FindStringSubmatch(string(text))
	if len(matches) == 0 {
		return fmt.Errorf("could not parse roll: %s", text)
	}

	for i, name := range rollRegex.SubexpNames() {
		switch name {
		case "dice":
			dice, err := strconv.Atoi(matches[i])
			if err != nil {
				return err
			}

			x.Dice = uint32(dice)
		case "die":
			die, err := strconv.Atoi(matches[i])
			if err != nil {
				return err
			}

			x.Die = uint32(die)
		case "keepScheme":
			if x.Keep == nil {
				x.Keep = &Keep{}
			}

			switch matches[i] {
			case "kh":
				x.Keep.Lowest = false
			case "kl":
				x.Keep.Lowest = true
			default:
				return fmt.Errorf("invalid keep scheme: %s", matches[i])
			}
		case "keepCount":
			if x.Keep == nil {
				x.Keep = &Keep{}
			}

			keep, err := strconv.Atoi(matches[i])
			if err != nil {
				return err
			}

			x.Keep.Count = uint32(keep)
		case "reroll":
			rrmatches := rerollRegex.FindAllStringSubmatch(string(matches[i]), -1)
			for _, rrmatch := range rrmatches {
				reroll := &Reroll{}
				for j, name := range rerollRegex.SubexpNames() {
					switch name {
					case "rerollDie":
						die, err := strconv.Atoi(rrmatch[j])
						if err != nil {
							return err
						}

						reroll.Die = uint32(die)
					case "rerollScheme":
						switch rrmatch[j] {
						case "rr":
							reroll.Once = false
						case "ro":
							reroll.Once = true
						}
					}
				}
				x.Reroll = append(x.Reroll, reroll)
			}
		case "bonus":
			bonus, err := strconv.Atoi(matches[i])
			if err != nil {
				return err
			}

			x.Bonus = uint32(bonus)
		}
	}

	return nil
}

func (x *Roll) MarshalText() (text []byte, err error) {
	str := fmt.Sprintf(rollFmt, x.Dice, x.Die)

	if x.GetKeep() != nil {
		keepScheme := "kh"
		if x.GetKeep().GetLowest() {
			keepScheme = "kl"
		}

		str += fmt.Sprintf(rollKeepFmt, keepScheme, x.GetKeep().GetCount())
	}

	for _, rr := range x.Reroll {
		rerollScheme := "rr"
		if rr.GetOnce() {
			rerollScheme = "ro"
		}

		str += fmt.Sprintf(rollRerollFmt, rerollScheme, rr.GetDie())
	}

	if x.GetBonus() != 0 {
		str += fmt.Sprintf(rollBonusFmt, x.GetBonus())
	}

	return []byte(str), nil
}
