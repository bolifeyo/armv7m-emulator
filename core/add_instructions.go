package core

import "fmt"

/* ADD (register)
 * ARM ARM A7.7.4
 * Encoding T1 */
type AddRegT1 InstrFields

func AddReg16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := RegIndex(raw_instr & 0x7)
	Rn := RegIndex((raw_instr >> 3) & 0x7)
	Rm := RegIndex((raw_instr >> 6) & 0x7)

	return AddRegT1{Rd: Rd, Rm: Rm, Rn: Rn, Imm: 0, setflags: NOT_IT}
}

func (instr AddRegT1) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegT1) String() string {
	return fmt.Sprintf("adds %s, %s, %s", instr.Rd, instr.Rn, instr.Rm)
}

/* ADD (SP plus register)
 * ARM ARM A7.7.6
 * Encoding T1 */
type AddRegSPT1 InstrFields

func AddRegSP16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	rdm := uint8(raw_instr & 0x7)
	DM := uint8((raw_instr >> 7) & 0x1)

	Rdm := RegIndex((DM << 3) | rdm)

	return AddRegSPT1{Rd: Rdm, Rm: Rdm, Rn: SP, Imm: 0, setflags: NEVER}
}

func (instr AddRegSPT1) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegSPT1) String() string {
	return fmt.Sprintf("add %s, sp, %s", instr.Rd, instr.Rd)
}

/* ADD (SP plus register)
 * ARM ARM A7.7.6
 * Encoding T2 */
type AddRegSPT2 InstrFields

func AddRegSP16T2(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rm := RegIndex((raw_instr >> 3) & 0xf)

	if Rm == SP {
		return AddRegSP16T1(instr)
	}

	return AddRegSPT2{Rd: SP, Rm: Rm, Rn: SP, Imm: 0, setflags: NEVER}
}

func (instr AddRegSPT2) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegSPT2) String() string {
	return fmt.Sprintf("add sp, %s", instr.Rm)
}
