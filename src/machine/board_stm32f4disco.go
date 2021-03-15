// +build stm32f4disco

package machine

import (
	"device/stm32"
	"runtime/interrupt"
)

const (
	LED         = LED_BUILTIN
	LED1        = LED_GREEN
	LED2        = LED_ORANGE
	LED3        = LED_RED
	LED4        = LED_BLUE
	LED_BUILTIN = LED_GREEN
	LED_GREEN   = PD12
	LED_ORANGE  = PD13
	LED_RED     = PD14
	LED_BLUE    = PD15
)

// UART pins
const (
	UART_TX_PIN = PA2
	UART_RX_PIN = PA3
)

var (
	UART0 = UART{
		Buffer:          NewRingBuffer(),
		Bus:             stm32.USART2,
		AltFuncSelector: AF7_USART1_2_3,
	}
	UART1 = &UART0
)

// set up RX IRQ handler. Follow similar pattern for other UARTx instances
func init() {
	UART0.Interrupt = interrupt.New(stm32.IRQ_USART2, UART0.handleInterrupt)
}

// SPI pins
const (
	SPI1_SCK_PIN = PA5
	SPI1_SDI_PIN = PA6
	SPI1_SDO_PIN = PA7
	SPI0_SCK_PIN = SPI1_SCK_PIN
	SPI0_SDI_PIN = SPI1_SDI_PIN
	SPI0_SDO_PIN = SPI1_SDO_PIN
)

// MEMs accelerometer
const (
	MEMS_ACCEL_CS   = PE3
	MEMS_ACCEL_INT1 = PE0
	MEMS_ACCEL_INT2 = PE1
)

// Since the first interface is named SPI1, both SPI0 and SPI1 refer to SPI1.
// TODO: implement SPI2 and SPI3.
var (
	SPI0 = SPI{
		Bus:             stm32.SPI1,
		AltFuncSelector: AF5_SPI1_SPI2,
	}
	SPI1 = &SPI0
)

const (
	I2C0_SCL_PIN = PB6
	I2C0_SDA_PIN = PB9
)

var (
	I2C0 = I2C{
		Bus:             stm32.I2C1,
		AltFuncSelector: AF4_I2C1_2_3,
	}
)
