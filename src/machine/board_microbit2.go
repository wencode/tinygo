// +build microbit2

package machine

// The micro:bit does not have a 32kHz crystal on board.
const HasLowFrequencyCrystal = false

// Buttons on the micro:bit (A, B and Logo)
const (
	BUTTON     Pin = BUTTONA
	BUTTONA    Pin = P0_14
	BUTTONB    Pin = P0_23
	LOGO_TOUCH Pin = P1_04
)

// UART pins
const (
	UART_TX_PIN Pin = P0_06
	UART_RX_PIN Pin = P1_08
)

// ADC pins
const (
	ADC0 Pin = 3 // todo: need confirm
	ADC1 Pin = 2 // todo: need confirm
	ADC2 Pin = 1 // todo: need confirm
)

// I2C pins
const (
	SDA_PIN Pin = P1_00
	SCL_PIN Pin = P0_26
)

// SPI pins
const (
	SPI0_SCK_PIN Pin = P0_17
	SPI0_SDO_PIN Pin = P0_01 // SPI MISO PIN
	SPI0_SDI_PIN Pin = P0_13 // SPI MOSI PIN
)

// GPIO/Analog pins
const (
	P0  Pin = P0_02
	P1  Pin = P0_03
	P2  Pin = P0_04
	P3  Pin = P0_31
	P4  Pin = P0_28
	P5  Pin = P0_14
	P6  Pin = P1_05
	P7  Pin = P0_11
	P8  Pin = P0_10
	P9  Pin = P0_09
	P10 Pin = P0_30
	P11 Pin = P0_23
	P12 Pin = P0_12
	P13 Pin = P0_17
	P14 Pin = P0_01
	P15 Pin = P0_13
	P16 Pin = P1_02
	P19 Pin = P0_26
	P20 Pin = P1_00
)

// LED matrix pins
const (
	LED_ROW_1 Pin = P0_21
	LED_ROW_2 Pin = P0_22
	LED_ROW_3 Pin = P0_15
	LED_ROW_4 Pin = P0_24
	LED_ROW_5 Pin = P0_19
	LED_COL_1 Pin = P0_28
	LED_COL_2 Pin = P0_11
	LED_COL_3 Pin = P0_31
	LED_COL_4 Pin = P1_05
	LED_COL_5 Pin = P0_30
)

// Audio
const (
	RUN_MIC_PIN Pin = P0_20
	MIC_IN_PIN  Pin = P0_05
	SPEAKER_PIN Pin = P0_00
)
