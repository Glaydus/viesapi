// Package access provides Modbus register definitions for Huawei SmartLogger devices.
//
// # Supported Devices
//
// This package contains complete Modbus register maps for:
//   - Huawei SmartLogger 3000 monitoring device
//   - Inverters connected to SmartLogger 3000
//
// # Huawei SmartLogger 3000
//
// SmartLogger 3000 is a monitoring and management device for Huawei photovoltaic installations.
// It communicates with inverters and other devices, aggregating data and providing access via Modbus TCP.
//
// System Time (40000-40009):
//   - dateTime (40000-40001): Date and time (epoch seconds)
//   - localTime (40009-40010): Local time (epoch seconds)
//
// Power Measurements (40521-40544):
//   - inputPower (40521-40522): DC input power in watts (gain: 1000)
//   - co2Reduction (40523-40524): CO2 reduction in kg (gain: 10)
//   - activePower (40525-40526): AC active power in watts (gain: 1000)
//   - powerFactor (40532): Power factor (gain: 1000)
//   - reactivePower (40544-40545): Reactive power in VAr (gain: 1000)
//
// Energy Counters (40560-40564):
//   - accumulatedEnergyYield (40560-40561): Total energy produced in kWh (gain: 10)
//   - dailyEnergyYield (40562-40563): Daily energy produced in kWh (gain: 10)
//   - dailyPowerDuration (40564-40565): Daily operating time in minutes (gain: 10)
//
// Phase Currents (40572-40574):
//   - phaseACurrent (40572): Phase A current in amperes
//   - phaseBCurrent (40573): Phase B current in amperes
//   - phaseCCurrent (40574): Phase C current in amperes
//
// Identification (40713-40736):
//   - esn (40713-40722): Device serial number
//   - deviceAccessStatus (40736): Device access status
//
// Alarms (50000-50001):
//   - alarm1 (50000): Primary alarms with error codes - scheduling issues
//   - alarm2 (50001): Secondary alarms with error codes - MCB, cubicle, power, license issues
//
// Each alarm bit includes an associated error code for detailed fault identification.
//
// Device Status (65534):
//   - deviceStatus (65534): Current device status with code mapping
//
// # SmartLogger - Connected Inverters
//
// Inverter registers connected to SmartLogger are accessible through relative addressing.
// Each inverter has its own set of registers starting from address 0.
//
// Power Measurements (0-8):
//   - activePower (0-1): Active power in watts (gain: 1000)
//   - reactivePower (2-3): Reactive power in VAr (gain: 1000)
//   - inputDc (4): DC input current in amperes (gain: 100)
//   - inputPower (5-6): DC input power in watts (gain: 1000)
//   - insulationResistance (7): Insulation resistance in kΩ (gain: 1000)
//   - powerFactor (8): Power factor (gain: 1000)
//
// Status and Temperature (9-11):
//   - deviceStatus (9): Device status with code mapping
//   - temperature (11): Temperature in °C (gain: 10)
//
// Fault Codes (12-18):
//   - majorFaultCode (12-13): Major fault code
//   - minorFaultCode (14-15): Minor fault code
//   - warningCode (16-17): Warning code
//
// # Data Types
//
// The following data types are used in register definitions:
//   - U16: Unsigned 16-bit integer (1 register)
//   - I16: Signed 16-bit integer (1 register)
//   - U32: Unsigned 32-bit integer (2 registers)
//   - I32: Signed 32-bit integer (2 registers)
//   - Str: ASCII string (multiple registers)
//   - Bitfield16: 16-bit bitfield (1 register)
//   - Strs: Array of strings (for merged properties)
//
// # Gain Values
//
// Gain represents the divisor for converting raw register values to actual values.
// For example, with gain=10 and raw value=235, the actual value is 23.5.
// With gain=1000 and raw value=5000, the actual value is 5.0.
//
// # Bitfield Error Codes
//
// For alarm registers (alarm1, alarm2), each bit definition includes not only
// the alarm description (On field) but also an associated error code (Code field).
// These codes are returned as part of an Alarm structure containing both the alarm
// name and its numeric code for detailed fault identification.
//
// # Codes Mapping
//
// The Codes field provides a mapping from numeric register values to human-readable strings.
// This is used for registers that represent enumerated states or status codes.
// When a register value matches a key in the Codes map, the corresponding string is returned.
// If no match is found, a formatted code string is returned (e.g., "Code: 0x123").
//
// Example usage in deviceStatus register (65534):
//   - Raw value 0xB000 returns "Disconnection"
//   - Raw value 0xB001 returns "Online"
//   - Raw value 0xFFFF (not in map) returns "Code: 0xFFFF"
//
// # Inverter Addressing
//
// SmartLogger manages multiple inverters. Access to registers of a specific inverter requires
// appropriate addressing through the Modbus protocol (typically via slave ID or address offset).
//
// For detailed register documentation, see smartlogger-modbus-definitions.md.
package access

import "installation/pkg/defs"

// smartDefaultProps defines the default set of properties to read from SmartLogger 3000.
// These properties provide essential operational data including status, power output,
// and energy production information.
var smartDefaultProps = []string{
	"deviceStatus", "accumulatedEnergyYield", "dailyEnergyYield", "dailyPowerDuration",
	"activePower", "inputPower", "reactivePower", "co2Reduction",
}

// smart3000Props contains the complete Modbus register map for Huawei SmartLogger 3000.
// Each property maps to specific Modbus registers with type information, scaling factors,
// and for bitfields, the meaning of each bit position.
//
// Register address ranges:
//   - 40000-40009: System time
//   - 40521-40574: Power measurements and phase currents
//   - 40713-40736: Device identification and access status
//   - 50000-50001: Alarms with error codes
//   - 65534: Device status
var smart3000Props = map[string]defs.PropDef{
	"dateTime":      {Addr: 40000, Type: "U32", Gain: 1, Quantity: 2},
	"localTime":     {Addr: 40009, Type: "U32", Gain: 1, Quantity: 2},
	"inputPower":    {Addr: 40521, Type: "U32", Gain: 1000, Quantity: 2},
	"co2Reduction":  {Addr: 40523, Type: "U32", Gain: 10, Quantity: 2},
	"activePower":   {Addr: 40525, Type: "I32", Gain: 1000, Quantity: 2},
	"powerFactor":   {Addr: 40532, Type: "I16", Gain: 1000, Quantity: 1},
	"reactivePower": {Addr: 40544, Type: "I32", Gain: 1000, Quantity: 2},

	"accumulatedEnergyYield": {Addr: 40560, Type: "U32", Gain: 10, Quantity: 2},
	"dailyEnergyYield":       {Addr: 40562, Type: "U32", Gain: 10, Quantity: 2},
	"dailyPowerDuration":     {Addr: 40564, Type: "U32", Gain: 10, Quantity: 2},
	"phaseACurrent":          {Addr: 40572, Type: "I16", Gain: 1, Quantity: 1},
	"phaseBCurrent":          {Addr: 40573, Type: "I16", Gain: 1, Quantity: 1},
	"phaseCCurrent":          {Addr: 40574, Type: "I16", Gain: 1, Quantity: 1},
	"esn":                    {Addr: 40713, Type: "Str", Gain: 1, Quantity: 10},
	"deviceAccessStatus":     {Addr: 40736, Type: "U16", Gain: 1, Quantity: 1},
	"alarm1": {Addr: 50000, Type: "Bitfield16", Gain: 1, Quantity: 1, Bits: defs.Bits{
		0b0000_0000_0000_1000: {On: "Abnormal Active Schedule", Code: 1100},
		0b0000_1000_0000_0000: {On: "Abnormal Reactive Schedule", Code: 1101},
	}},
	"alarm2": {Addr: 50001, Type: "Bitfield16", Gain: 1, Quantity: 1, Bits: defs.Bits{
		0b0000_0000_0000_0001: {On: "MCB Disconnect", Code: 1103},
		0b0000_0000_0000_0010: {On: "Abnormal Cubicle", Code: 1104},
		0b0000_0000_0000_0100: {On: "Device Address Conflict", Code: 1105},
		0b0000_0000_0000_1000: {On: "AC SPD fault", Code: 1106},
		0b0010_0000_0000_0000: {On: "24V power failure", Code: 1115},
		0b0100_0000_0000_0000: {On: "License Expired", Code: 1119},
	}},
	"alarms": {
		Merge: []string{"alarm1", "alarm2"},
		Type:  "Strs",
	},
	"deviceStatus": {Addr: 65534, Type: "U16", Gain: 1, Quantity: 1, Codes: defs.Codes{
		0xB000: "Disconnection",
		0xB001: "Online",
	}},
}

// smartInverterProps contains the complete Modbus register map for inverters connected to SmartLogger 3000.
// Each property maps to specific Modbus registers with type information and scaling factors.
// These registers use relative addressing starting from 0 for each inverter.
//
// Register address ranges:
//   - 0-8: Power measurements and power factor
//   - 9-11: Device status and temperature
//   - 12-18: Fault and warning codes
var smartInverterProps = map[string]defs.PropDef{
	"activePower":          {Addr: 0, Type: "I32", Gain: 1000, Quantity: 2},
	"reactivePower":        {Addr: 2, Type: "I32", Gain: 1000, Quantity: 2},
	"inputDc":              {Addr: 4, Type: "I16", Gain: 100, Quantity: 1},
	"inputPower":           {Addr: 5, Type: "U32", Gain: 1000, Quantity: 2},
	"insulationResistance": {Addr: 7, Type: "U16", Gain: 1000, Quantity: 1},
	"powerFactor":          {Addr: 8, Type: "I16", Gain: 1000, Quantity: 1},
	"deviceStatus": {Addr: 9, Type: "U16", Gain: 1, Quantity: 1, Codes: defs.Codes{
		0xB000: "Communication interrupt",
		0xB001: "Online",
		0xC000: "Uploading",
	}},
	"temperature":    {Addr: 11, Type: "I16", Gain: 10, Quantity: 1},
	"majorFaultCode": {Addr: 12, Type: "U32", Gain: 1, Quantity: 2},
	"minorFaultCode": {Addr: 14, Type: "U32", Gain: 1, Quantity: 2},
	"warningCode":    {Addr: 16, Type: "U32", Gain: 1, Quantity: 2},
}
