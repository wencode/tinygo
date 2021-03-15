// +build stm32l5x2

package machine

// Peripheral abstraction layer for the stm32f407

import (
	"device/stm32"
)

func CPUFrequency() uint32 {
	return 110000000
}

//---------- UART related code

// Configure the UART.
func (uart *UART) configurePins(config UARTConfig) {
	if config.RX.getPort() == stm32.GPIOG || config.TX.getPort() == stm32.GPIOG {
		// Enable VDDIO2 power supply, which is an independant power supply for the PGx pins
		stm32.PWR.CR2.SetBits(stm32.PWR_CR2_IOSV)
	}

	// enable the alternate functions on the TX and RX pins
	config.TX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTTX}, uart.AltFuncSelector)
	config.RX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTRX}, uart.AltFuncSelector)
}

// UART baudrate calc based on the bus and clockspeed
// NOTE: keep this in sync with the runtime/runtime_stm32l5x2.go clock init code
func (uart *UART) getBaudRateDivisor(baudRate uint32) uint32 {
	return 256 * (CPUFrequency() / baudRate)
}

// Register names vary by ST processor, these are for STM L5
func (uart *UART) setRegisters() {
	uart.rxReg = &uart.Bus.RDR
	uart.txReg = &uart.Bus.TDR
	uart.statusReg = &uart.Bus.ISR
	uart.txEmptyFlag = stm32.USART_ISR_TXE
}
