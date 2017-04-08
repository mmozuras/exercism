package queenattack

import "errors"

const testVersion = 2

func CanQueenAttack(white, black string) (bool, error) {
	if !isValid(white) || !isValid(black) || white == black {
		return false, errors.New("Invalid position")
	}

	wf := int(white[0])
	wr := int(white[1])
	bf := int(black[0])
	br := int(black[1])

	return (wf == bf || wr == br || bf-wf == br-wr || bf-wf == -(br-wr)), nil
}

func isValid(pos string) bool {
	return len(pos) == 2 &&
		pos[0] >= 'a' && pos[0] <= 'h' &&
		pos[1] >= '1' && pos[1] <= '8'
}
