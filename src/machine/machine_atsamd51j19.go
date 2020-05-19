// +build sam,atsamd51,atsamd51j19

// Peripheral abstraction layer for the atsamd51.
//
// Datasheet:
// http://ww1.microchip.com/downloads/en/DeviceDoc/SAM_D5xE5x_Family_Data_Sheet_DS60001507F.pdf
//
package machine

import "device/sam"

const HSRAM_SIZE = 0x00030000

// This chip has five TCC peripherals, which have PWM as one feature.
var (
	PWM0 = (*PWM)(sam.TCC0)
	PWM1 = (*PWM)(sam.TCC1)
	PWM2 = (*PWM)(sam.TCC2)
	PWM3 = (*PWM)(sam.TCC3)
	PWM4 = (*PWM)(sam.TCC4)
)

func (pwm *PWM) configureClock() {
	// Turn on timer clocks used for PWM and use generic clock generator 0.
	switch pwm.timer() {
	case sam.TCC0:
		sam.MCLK.APBBMASK.SetBits(sam.MCLK_APBBMASK_TCC0_)
		sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC0].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) | sam.GCLK_PCHCTRL_CHEN)
	case sam.TCC1:
		sam.MCLK.APBBMASK.SetBits(sam.MCLK_APBBMASK_TCC1_)
		sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC1].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) | sam.GCLK_PCHCTRL_CHEN)
	case sam.TCC2:
		sam.MCLK.APBCMASK.SetBits(sam.MCLK_APBCMASK_TCC2_)
		sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC2].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) | sam.GCLK_PCHCTRL_CHEN)
	case sam.TCC3:
		sam.MCLK.APBCMASK.SetBits(sam.MCLK_APBCMASK_TCC3_)
		sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC3].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) | sam.GCLK_PCHCTRL_CHEN)
	case sam.TCC4:
		sam.MCLK.APBDMASK.SetBits(sam.MCLK_APBDMASK_TCC4_)
		sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC4].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) | sam.GCLK_PCHCTRL_CHEN)
	}
}

func (pwm *PWM) timerNum() uint8 {
	switch pwm.timer() {
	case sam.TCC0:
		return 0
	case sam.TCC1:
		return 1
	case sam.TCC2:
		return 2
	case sam.TCC3:
		return 3
	case sam.TCC4:
		return 4
	default:
		return 0x0f // should not happen
	}
}
