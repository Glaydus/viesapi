// Package access provides Modbus register definitions for supported inverter models.
//
// # Supported Devices
//
// This package contains complete Modbus register maps for:
//   - Huawei SUN2000 series solar inverters
//   - SolarEdge solar inverters
//
// # Huawei SUN2000 Inverter
//
// The SUN2000 series uses Modbus TCP/RTU protocol with the following register map:
//
// Device Information (30000-30082):
//   - modelName (30000-30014): Device model name
//   - sn (30015-30024): Serial number
//   - pn (30025-30034): Part number
//   - firmwareVer (30035-30049): Firmware version
//   - softwareVer (30050-30064): Software version
//   - modelId (30070): Model identifier
//   - nofPvStrings (30071): Number of PV strings
//   - nofMppTrackers (30072): Number of MPP trackers
//   - ratedPower (30073-30074): Rated power in watts (gain: 1000)
//   - pmax (30075-30076): Maximum active power (gain: 1000)
//   - smax (30077-30078): Maximum apparent power (gain: 1000)
//   - qmaxFed (30079-30080): Maximum reactive power fed to grid (gain: 1000)
//   - qmaxAbsorbed (30081-30082): Maximum reactive power absorbed from grid (gain: 1000)
//
// Device States (32000-32004):
//   - state1 (32000): Primary device state bitfield
//     Bit 0: standby, Bit 1: grid connected, Bit 2: grid-connected normally
//     Bit 3: derating due to power rationing, Bit 4: derating due to internal causes
//     Bit 5: normal stop, Bit 6: stop due to faults, Bit 7: stop due to power rationing
//     Bit 8: shutdown, Bit 9: spot check
//   - state2 (32002): Secondary device state bitfield
//     Bit 0: unlocked/locked, Bit 1: PV connected/disconnected
//     Bit 2: DSP data collection on/off
//   - state3 (32003-32004): Grid connection state
//     Bit 0: off-grid/on-grid, Bit 1: off-grid switch enabled/disabled
//
// Alarms (32008-32010):
//   - alarm1 (32008): Primary alarms with error codes - string/grid voltage issues, arc faults, grid loss
//   - alarm2 (32009): Secondary alarms with error codes - grounding, temperature, device faults, battery issues
//   - alarm3 (32010): Tertiary alarms with error codes - optimizer, fan, DC protection, CT wiring issues
//
// Each alarm bit includes an associated error code for detailed fault identification.
//
// PV String Measurements (32016-32062):
//   - pvXVoltage: String voltage in volts (gain: 10), supports up to 24 strings
//   - pvXCurrent: String current in amps (gain: 100), supports up to 24 strings
//
// Grid Measurements (32064-32090):
//   - inputPower (32064-32065): Total DC input power (gain: 1000)
//   - Grid voltages (32066-32071): Line and phase voltages (gain: 10)
//   - Grid currents (32072-32077): Phase currents (gain: 1000)
//   - activePower (32080-32081): AC active power output (gain: 1000)
//   - reactivePower (32082-32083): AC reactive power (gain: 1000)
//   - powerFactor (32084): Power factor (gain: 1000)
//   - gridFrequency (32085): Grid frequency in Hz (gain: 100)
//   - efficiency (32086): Conversion efficiency percentage (gain: 100)
//   - internalTemperature (32087): Internal temperature in °C (gain: 10)
//   - deviceStatus (32089): Current device status code
//   - faultCode (32090): Active fault code
//
// Energy Counters (32091-32119):
//   - startupTime (32091-32092): Last startup timestamp (epoch seconds, local time)
//   - shutdownTime (32093-32094): Last shutdown timestamp (epoch seconds, local time)
//   - accumulatedEnergyYield (32106-32107): Total energy produced in kWh (gain: 100)
//   - totalInputPower (32108-32109): Total input power in kW (gain: 100)
//   - dailyEnergyYield (32114-32115): Daily energy produced in kWh (gain: 100)
//   - monthlyEnergyYield (32116-32117): Monthly energy produced in kWh (gain: 100)
//   - yearlyEnergyYield (32118-32119): Yearly energy produced in kWh (gain: 100)
//
// Power Control (35300-35307):
//   - activeAdjustmentMode (35300): 0=percentage, 1=fixed value
//   - reactiveAdjustmentMode (35304): Reactive power adjustment mode
//
// Power Grid Scheduling (40000-42020):
//   - systemTime (40000-40001): System time (epoch seconds, local time)
//   - pgsReactivePowerCompensationPf (40122): Reactive power compensation PF (gain: 1000)
//   - pgsActivePowerDerating (40125): Active power derating percentage (gain: 10)
//   - Various Q-U, PF-U, and cosφ-P/Pn characteristic curves
//
// # SolarEdge Inverter
//
// SolarEdge inverters use SunSpec Modbus protocol with scaled values:
//
// Device Information (40020-40069):
//   - modelName (40020-40035): Device model name
//   - pn (40044-40051): Part number
//   - sn (40052-40067): Serial number
//   - modelId (40069): Model identifier
//
// Measurements (40071-40107):
//   - Currents (40071-40074): Total and phase currents with scale factor at 40075
//   - Voltages (40076-40081): Line and phase voltages with scale factor at 40082
//   - activePower (40083): Active power with scale factor at 40084
//   - gridFrequency (40085): Grid frequency with scale factor at 40086
//   - reactivePower (40089): Reactive power with scale factor at 40090
//   - powerFactor (40091): Power factor with scale factor at 40092
//   - accumulatedEnergyYield (40093-40094): Total energy with scale factor at 40095
//   - temperature (40103): Heatsink temperature with scale factor at 40106
//   - deviceStatus (40107): Operating state
//
// Note: SolarEdge uses ScaleAddr fields where the value = raw_value x 10^(scale_register_value)
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
//   - Bitfield32: 32-bit bitfield (2 registers)
//   - Bytes: Raw byte array (multiple registers)
//
// # Gain Values
//
// Gain represents the divisor for converting raw register values to actual values.
// For example, with gain=10 and raw value=235, the actual value is 23.5.
// With gain=1000 and raw value=5000, the actual value is 5.0.
//
// # Bitfield Error Codes
//
// For alarm registers (alarm1, alarm2, alarm3), each bit definition includes not only
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
// Example usage in deviceStatus register (32089):
//   - Raw value 0x0200 returns "On grid"
//   - Raw value 0x0201 returns "Grid connection: power limited"
//   - Raw value 0xFFFF (not in map) returns "Code: 0xFFFF"
//
// # Energy Meters
//
// For detailed documentation of NR30 series energy meter registers, see meterdef.go.
// Energy meters provide comprehensive three-phase power monitoring including:
//   - Real-time voltage, current, power measurements per phase
//   - Energy counters (imported, exported, reactive, apparent)
//   - Alarm and status monitoring
//   - Network and archiving status
package access

import "installation/pkg/defs"

// sun2000DefaultProps defines the default set of properties to read from Huawei SUN2000 inverters.
// These properties provide essential operational data including identification, status, power output,
// and alarm information.
var sun2000DefaultProps = []string{
	"modelName", "sn", "pn", "modelId", "deviceStatus", "dailyEnergyYield",
	"accumulatedEnergyYield", "activePower", "reactivePower", "efficiency", "dayPeekActivePower",
	"pv1Voltage", "pv1Current", "pgsReactivePowerCompensationPf", "pgsActivePowerDerating",
	"faultCode", "states", "alarms",
}

// solarEdgeDefaultProps defines the default set of properties to read from SolarEdge inverters.
// These properties provide essential operational data using the SunSpec Modbus protocol.
var solarEdgeDefaultProps = []string{
	"modelName", "sn", "pn", "modelId", "deviceStatus", "accumulatedEnergyYield", "activePower", "reactivePower", "powerFactor", "temperature",
}

// sun2000Props contains the complete Modbus register map for Huawei SUN2000 inverters.
// Each property maps to specific Modbus registers with type information, scaling factors,
// and for bitfields, the meaning of each bit position.
//
// Register address ranges:
//   - 30000-30082: Device identification and capabilities
//   - 32000-32119: Real-time measurements and status
//   - 35300-35307: Power adjustment controls
//   - 37113-37518: Power meter and optimizer data
//   - 40000-40201: System configuration and grid scheduling
//   - 42000-42020: Grid code and power change parameters
//   - 43006: Time zone configuration
var sun2000Props = map[string]defs.PropDef{
	"modelName":      {Type: "Str", Gain: 1, Addr: 30000, Quantity: 15},
	"sn":             {Type: "Str", Gain: 1, Addr: 30015, Quantity: 10},
	"pn":             {Type: "Str", Gain: 1, Addr: 30025, Quantity: 10},
	"firmwareVer":    {Type: "Str", Gain: 1, Addr: 30035, Quantity: 15},
	"softwareVer":    {Type: "Str", Gain: 1, Addr: 30050, Quantity: 15},
	"modelId":        {Type: "U16", Gain: 1, Addr: 30070, Quantity: 1},
	"nofPvStrings":   {Type: "U16", Gain: 1, Addr: 30071, Quantity: 1},
	"nofMppTrackers": {Type: "U16", Gain: 1, Addr: 30072, Quantity: 1},
	"ratedPower":     {Type: "U32", Gain: 1000, Addr: 30073, Quantity: 2},
	"pmax":           {Type: "U32", Gain: 1000, Addr: 30075, Quantity: 2}, // Maximum active power
	"smax":           {Type: "U32", Gain: 1000, Addr: 30077, Quantity: 2}, // Maximum apparent power
	"qmaxFed":        {Type: "I32", Gain: 1000, Addr: 30079, Quantity: 2}, // Maximum reactive power (Qmax, fed to the power grid)
	"qmaxAbsorbed":   {Type: "I32", Gain: 1000, Addr: 30081, Quantity: 2}, // Maximum reactive power (Qmax, absorbed from the power grid)
	"state1": {Type: "Bitfield16", Gain: 1, Addr: 32000, Quantity: 1, Bits: defs.Bits{
		0b00_0000_0001: {On: "standby", Off: ""},
		0b00_0000_0010: {On: "grid connected", Off: ""},
		0b00_0000_0100: {On: "grid-connected normally", Off: ""},
		0b00_0000_1000: {On: "grid connection with derating due to power rationing", Off: ""},
		0b00_0001_0000: {On: "grid connection with derating due to internal causes of the solar inverter", Off: ""},
		0b00_0010_0000: {On: "normal stop", Off: ""},
		0b00_0100_0000: {On: "stop due to faults", Off: ""},
		0b00_1000_0000: {On: "stop due to power rationing", Off: ""},
		0b01_0000_0000: {On: "shutdown", Off: ""},
		0b10_0000_0000: {On: "spot check", Off: ""},
	}},
	"state2": {Type: "Bitfield16", Gain: 1, Addr: 32002, Quantity: 1, Bits: defs.Bits{
		0b001: {On: "unlocked", Off: "locked"},
		0b010: {On: "PV connected", Off: "PV disconnected"},
		0b100: {On: "DSP data collection on", Off: "DSP data collection off"},
	}},
	"state3": {Type: "Bitfield32", Gain: 1, Addr: 32003, Quantity: 2, Bits: defs.Bits{
		0b01: {On: "off-grid", Off: "on-grid"},
		0b10: {On: "off-grid switch enabled", Off: "off-grid switch disabled"},
	}},
	"states": {
		Merge: []string{"state1", "state2", "state3"},
		Type:  "Strs",
	},
	"alarm1": {Type: "Bitfield16", Gain: 1, Addr: 32008, Quantity: 1, Bits: defs.Bits{
		0b0000_0000_0000_0001: {On: "High String Input Voltage", Code: 2001},
		0b0000_0000_0000_0010: {On: "DC Arc Fault", Code: 2002},
		0b0000_0000_0000_0100: {On: "String Reverse Connection", Code: 2011},
		0b0000_0000_0000_1000: {On: "String Current Backfeed", Code: 2012},
		0b0000_0000_0001_0000: {On: "Abnormal String Power", Code: 2013},
		0b0000_0000_0010_0000: {On: "AFCI Self-Check Fail.", Code: 2021},
		0b0000_0000_0100_0000: {On: "Phase Wire Short-Circuited to PE", Code: 2031},
		0b0000_0000_1000_0000: {On: "Grid Loss", Code: 2032},
		0b0000_0001_0000_0000: {On: "Grid Undervoltage", Code: 2033},
		0b0000_0010_0000_0000: {On: "Grid Overvoltage", Code: 2034},
		0b0000_0100_0000_0000: {On: "Grid Volt. Imbalance", Code: 2035},
		0b0000_1000_0000_0000: {On: "Grid Overfrequency", Code: 2036},
		0b0001_0000_0000_0000: {On: "Grid Underfrequency", Code: 2037},
		0b0010_0000_0000_0000: {On: "Unstable Grid Frequency", Code: 2038},
		0b0100_0000_0000_0000: {On: "Output Overcurrent", Code: 2039},
		0b1000_0000_0000_0000: {On: "Output DC Component Overhigh", Code: 2040},
	}},
	"alarm2": {Type: "Bitfield16", Gain: 1, Addr: 32009, Quantity: 1, Bits: defs.Bits{
		0b0000_0000_0000_0001: {On: "Abnormal Residual Current", Code: 2051},
		0b0000_0000_0000_0010: {On: "Abnormal Grounding", Code: 2061},
		0b0000_0000_0000_0100: {On: "Low Insulation Resistance", Code: 2062},
		0b0000_0000_0000_1000: {On: "Overtemperature", Code: 2063},
		0b0000_0000_0001_0000: {On: "Device Fault", Code: 2064},
		0b0000_0000_0010_0000: {On: "Upgrade Failed or Version Mismatch", Code: 2065},
		0b0000_0000_0100_0000: {On: "License Expired", Code: 2066},
		0b0000_0000_1000_0000: {On: "Faulty Monitoring Unit", Code: 61440},
		0b0000_0001_0000_0000: {On: "Faulty Power Collector", Code: 2067},
		0b0000_0010_0000_0000: {On: "Battery abnormal", Code: 2068},
		0b0000_0100_0000_0000: {On: "Active Islanding", Code: 2070},
		0b0000_1000_0000_0000: {On: "Passive Islanding", Code: 2071},
		0b0001_0000_0000_0000: {On: "Transient AC Overvoltage", Code: 2072},
		0b0010_0000_0000_0000: {On: "Peripheral port short circuit", Code: 2075},
		0b0100_0000_0000_0000: {On: "Churn output overload", Code: 2077},
		0b1000_0000_0000_0000: {On: "Abnormal PV module configuration", Code: 2080},
	}},
	"alarm3": {Type: "Bitfield16", Gain: 1, Addr: 32010, Quantity: 1, Bits: defs.Bits{
		0b0000_0000_0000_0001: {On: "Optimizer fault", Code: 2081},
		0b0000_0000_0000_0010: {On: "Built-in PID operation abnormal", Code: 2085},
		0b0000_0000_0000_0100: {On: "High input string voltage to ground.", Code: 2014},
		0b0000_0000_0000_1000: {On: "External Fan Abnormal", Code: 2086},
		0b0000_0000_0001_0000: {On: "Battery Reverse Connection", Code: 2069},
		0b0000_0000_0010_0000: {On: "On-grid/Off-grid controller abnormal", Code: 2082},
		0b0000_0000_0100_0000: {On: "PV String Loss", Code: 2015},
		0b0000_0000_1000_0000: {On: "Internal Fan Abnormal", Code: 2087},
		0b0000_0001_0000_0000: {On: "DC Protection Unit Abnormal", Code: 2088},
		0b0000_0010_0000_0000: {On: "EL Unit Abnormal", Code: 2089},
		0b0000_0100_0000_0000: {On: "Active Adjustment Instruction Abnormal", Code: 2090},
		0b0000_1000_0000_0000: {On: "Reactive Adjustment Instruction Abnormal", Code: 2091},
		0b0001_0000_0000_0000: {On: "CT Wiring Abnormal", Code: 2092},
		0b0010_0000_0000_0000: {On: "DC Arc Fault(ADMC Alarm to be clear manually)", Code: 2003},
	}},
	"alarms": {
		Merge: []string{"alarm1", "alarm2", "alarm3"},
		Type:  "Strs",
	},
	"pv1Voltage":                    {Type: "I16", Gain: 10, Addr: 32016, Quantity: 1},
	"pv1Current":                    {Type: "I16", Gain: 100, Addr: 32017, Quantity: 1},
	"pv2Voltage":                    {Type: "I16", Gain: 10, Addr: 32018, Quantity: 1},
	"pv2Current":                    {Type: "I16", Gain: 100, Addr: 32019, Quantity: 1},
	"pv3Voltage":                    {Type: "I16", Gain: 10, Addr: 32020, Quantity: 1},
	"pv3Current":                    {Type: "I16", Gain: 100, Addr: 32021, Quantity: 1},
	"pv4Voltage":                    {Type: "I16", Gain: 10, Addr: 32022, Quantity: 1},
	"pv4Current":                    {Type: "I16", Gain: 100, Addr: 32023, Quantity: 1},
	"pv5Voltage":                    {Type: "I16", Gain: 10, Addr: 32024, Quantity: 1},
	"pv5Current":                    {Type: "I16", Gain: 100, Addr: 32025, Quantity: 1},
	"pv6Voltage":                    {Type: "I16", Gain: 10, Addr: 32026, Quantity: 1},
	"pv6Current":                    {Type: "I16", Gain: 100, Addr: 32027, Quantity: 1},
	"pv7Voltage":                    {Type: "I16", Gain: 10, Addr: 32028, Quantity: 1},
	"pv7Current":                    {Type: "I16", Gain: 100, Addr: 32029, Quantity: 1},
	"pv8Voltage":                    {Type: "I16", Gain: 10, Addr: 32030, Quantity: 1},
	"pv8Current":                    {Type: "I16", Gain: 100, Addr: 32031, Quantity: 1},
	"pv9Voltage":                    {Type: "I16", Gain: 10, Addr: 32032, Quantity: 1},
	"pv9Current":                    {Type: "I16", Gain: 100, Addr: 32033, Quantity: 1},
	"pv10Voltage":                   {Type: "I16", Gain: 10, Addr: 32034, Quantity: 1},
	"pv10Current":                   {Type: "I16", Gain: 100, Addr: 32034, Quantity: 1},
	"pv11Voltage":                   {Type: "I16", Gain: 10, Addr: 32035, Quantity: 1},
	"pv11Current":                   {Type: "I16", Gain: 100, Addr: 32036, Quantity: 1},
	"pv12Voltage":                   {Type: "I16", Gain: 10, Addr: 32037, Quantity: 1},
	"pv12Current":                   {Type: "I16", Gain: 100, Addr: 32038, Quantity: 1},
	"pv13Voltage":                   {Type: "I16", Gain: 10, Addr: 32039, Quantity: 1},
	"pv13Current":                   {Type: "I16", Gain: 100, Addr: 32040, Quantity: 1},
	"pv14Voltage":                   {Type: "I16", Gain: 10, Addr: 32041, Quantity: 1},
	"pv14Current":                   {Type: "I16", Gain: 100, Addr: 32042, Quantity: 1},
	"pv15Voltage":                   {Type: "I16", Gain: 10, Addr: 32043, Quantity: 1},
	"pv15Current":                   {Type: "I16", Gain: 100, Addr: 32044, Quantity: 1},
	"pv16Voltage":                   {Type: "I16", Gain: 10, Addr: 32045, Quantity: 1},
	"pv16Current":                   {Type: "I16", Gain: 100, Addr: 32046, Quantity: 1},
	"pv17Voltage":                   {Type: "I16", Gain: 10, Addr: 32047, Quantity: 1},
	"pv17Current":                   {Type: "I16", Gain: 100, Addr: 32048, Quantity: 1},
	"pv18Voltage":                   {Type: "I16", Gain: 10, Addr: 32049, Quantity: 1},
	"pv18Current":                   {Type: "I16", Gain: 100, Addr: 32050, Quantity: 1},
	"pv19Voltage":                   {Type: "I16", Gain: 10, Addr: 32051, Quantity: 1},
	"pv19Current":                   {Type: "I16", Gain: 100, Addr: 32052, Quantity: 1},
	"pv20Voltage":                   {Type: "I16", Gain: 10, Addr: 32053, Quantity: 1},
	"pv20Current":                   {Type: "I16", Gain: 100, Addr: 32054, Quantity: 1},
	"pv21Voltage":                   {Type: "I16", Gain: 10, Addr: 32055, Quantity: 1},
	"pv21Current":                   {Type: "I16", Gain: 100, Addr: 32056, Quantity: 1},
	"pv22Voltage":                   {Type: "I16", Gain: 10, Addr: 32057, Quantity: 1},
	"pv22Current":                   {Type: "I16", Gain: 100, Addr: 32058, Quantity: 1},
	"pv23Voltage":                   {Type: "I16", Gain: 10, Addr: 32059, Quantity: 1},
	"pv23Current":                   {Type: "I16", Gain: 100, Addr: 32060, Quantity: 1},
	"pv24Voltage":                   {Type: "I16", Gain: 10, Addr: 32061, Quantity: 1},
	"pv24Current":                   {Type: "I16", Gain: 100, Addr: 32062, Quantity: 1},
	"inputPower":                    {Type: "I32", Gain: 1000, Addr: 32064, Quantity: 2},
	"gridVoltageOrLineVoltageAAndB": {Type: "U16", Gain: 10, Addr: 32066, Quantity: 1}, // Power grid voltage/Line voltage between phases A and B
	"lineVoltageBAndC":              {Type: "U16", Gain: 10, Addr: 32067, Quantity: 1}, // Line voltage between phases B and C
	"lineVoltageCAndA":              {Type: "U16", Gain: 10, Addr: 32068, Quantity: 1}, // Line voltage between phases B and C
	"phaseAVoltage":                 {Type: "U16", Gain: 10, Addr: 32069, Quantity: 1},
	"phaseBVoltage":                 {Type: "U16", Gain: 10, Addr: 32070, Quantity: 1},
	"phaseCVoltage":                 {Type: "U16", Gain: 10, Addr: 32071, Quantity: 1},
	"gridCurrentOrPhaseACurrent":    {Type: "I32", Gain: 1000, Addr: 32072, Quantity: 2}, // Power grid current/Phase A current
	"phaseBCurrent":                 {Type: "I32", Gain: 1000, Addr: 32074, Quantity: 2},
	"phaseCCurrent":                 {Type: "I32", Gain: 1000, Addr: 32076, Quantity: 2},
	"dayPeekActivePower":            {Type: "I32", Gain: 1000, Addr: 32078, Quantity: 2}, // Peak active power of current day
	"activePower":                   {Type: "I32", Gain: 1000, Addr: 32080, Quantity: 2},
	"reactivePower":                 {Type: "I32", Gain: 1000, Addr: 32082, Quantity: 2},
	"powerFactor":                   {Type: "I16", Gain: 1000, Addr: 32084, Quantity: 1},
	"gridFrequency":                 {Type: "U16", Gain: 100, Addr: 32085, Quantity: 1},
	"efficiency":                    {Type: "U16", Gain: 100, Addr: 32086, Quantity: 1},
	"internalTemperature":           {Type: "I16", Gain: 10, Addr: 32087, Quantity: 1},
	"insulationResistance":          {Type: "U16", Gain: 1000, Addr: 32088, Quantity: 1},
	"deviceStatus": {Type: "U16", Gain: 1, Addr: 32089, Quantity: 1, Codes: defs.Codes{
		0x0000: "Standby: initializing",
		0x0001: "Standby: detecting insulation resistance",
		0x0002: "Standby: detecting irradiation",
		0x0003: "Standby: grid detecting",
		0x0100: "Starting",
		0x0200: "On grid",
		0x0201: "Grid connection: power limited",
		0x0202: "Grid connection: self-derating",
		0x0203: "Off-grid Running",
		0xA000: "Idle: no irradiation",
		0xB000: "Communication interrupt",
		0xB001: "Online",
		0xC000: "Uploading",
	}},
	"faultCode":               {Type: "U16", Gain: 1, Addr: 32090, Quantity: 1},
	"startupTime":             {Type: "U32", Gain: 1, Addr: 32091, Quantity: 2}, // Epoch seconds, local time
	"shutdownTime":            {Type: "U32", Gain: 1, Addr: 32093, Quantity: 2}, // Epoch seconds, local time
	"accumulatedEnergyYield":  {Type: "U32", Gain: 100, Addr: 32106, Quantity: 2},
	"totalInputPower":         {Type: "U32", Gain: 100, Addr: 32108, Quantity: 2},
	"dailyEnergyYield":        {Type: "U32", Gain: 100, Addr: 32114, Quantity: 2},
	"monthlyEnergyYield":      {Type: "U32", Gain: 100, Addr: 32116, Quantity: 2},
	"yearlyEnergyYield":       {Type: "U32", Gain: 100, Addr: 32118, Quantity: 2},
	"activeRegulationState":   {Type: "Bytes", Gain: 1, Addr: 35300, Quantity: 4},
	"reactiveRegulationState": {Type: "Bytes", Gain: 1, Addr: 35304, Quantity: 4},
	"pmcActivePower":          {Type: "I32", Gain: 1, Addr: 37113, Quantity: 2}, // [Power meter collection] Active power*
	"nofTotalOptimizers":      {Type: "U16", Gain: 1, Addr: 37200, Quantity: 1}, // [Optimizer] Total number of optimizers*
	"nofOnlineOptimizers":     {Type: "U16", Gain: 1, Addr: 37201, Quantity: 1}, // [Optimizer] Number of online optimizers*
	"optimizerFeatureData":    {Type: "U16", Gain: 1, Addr: 37202, Quantity: 1},
	"inverterStatus": {Type: "U16", Gain: 1, Addr: 37518, Quantity: 1, Codes: defs.Codes{
		0: "offline",
		1: "online",
	}},
	"systemTime":                          {Type: "U32", Gain: 1, Addr: 40000, Quantity: 2}, // [946684800, 3155759999] Epoch seconds, local time
	"pgsQuCurveMode":                      {Type: "U16", Gain: 1, Addr: 40037, Quantity: 1}, // [Power grid scheduling] Q-U characteristic curve mode*
	"pgsQuDispatchTriggerPower":           {Type: "U16", Gain: 1, Addr: 40038, Quantity: 1},
	"pgsFixedActivePowerDerated":          {Type: "U16", Gain: 1, Addr: 40120, Quantity: 1},
	"pgsReactivePowerCompensationPf":      {Type: "I16", Gain: 1000, Addr: 40122, Quantity: 1}, // [Power grid scheduling] Reactive power compensation (PF)
	"pgsReactivePowerCompensationQs":      {Type: "I16", Gain: 1000, Addr: 40123, Quantity: 1}, // [Power grid scheduling] Reactive power compensation (Q/S)
	"pgsActivePowerDerating":              {Type: "U16", Gain: 10, Addr: 40125, Quantity: 1},   // [Power grid scheduling] Active power percentage derating (0.1%)
	"pgsFixedActivePowerDeratedW":         {Type: "U32", Gain: 1, Addr: 40126, Quantity: 2},    // [Power grid scheduling] Fixed active power derated (W)
	"pgsReactivePowerCompensationNightly": {Type: "I32", Gain: 1000, Addr: 40129, Quantity: 2}, // [Power grid scheduling] Reactive power compensation at night (kVar)
	"pgsCosFiPpnCurve":                    {Type: "Bytes", Gain: 1, Addr: 40133, Quantity: 21}, // [Power grid scheduling] cosφ-P/Pn characteristic curve
	"pgsQuCurve":                          {Type: "Bytes", Gain: 1, Addr: 40154, Quantity: 21}, // [Power grid scheduling] Q-U characteristic curve
	"pgsPfuCurve":                         {Type: "Bytes", Gain: 1, Addr: 40175, Quantity: 21}, // [Power grid scheduling] PF-U characteristic curve
	"pgsReactivePowerAdjustmentTime":      {Type: "U16", Gain: 1, Addr: 40196, Quantity: 1},    // [Power grid scheduling] Reactive power adjustment time
	"pgsQuPowPercToExitScheduling":        {Type: "U16", Gain: 1, Addr: 40198, Quantity: 1},    // [Power grid scheduling] Q-U power percentage to exit scheduling*
	"startup":                             {Type: "U16", Gain: 1, Addr: 40200, Quantity: 1},
	"shutdown":                            {Type: "U16", Gain: 1, Addr: 40201, Quantity: 1},
	"gridCode": {Type: "U16", Gain: 1, Addr: 42000, Quantity: 1, Codes: defs.Codes{
		303: "Poland-EN50549-1-LV",
		304: "Poland-EN50549-1-MV",
		305: "Poland-NC-RfG-LV",
	}},
	"pgsReactivePowerChangeGradient":      {Type: "U32", Gain: 1000, Addr: 42015, Quantity: 2}, // [Power grid scheduling] Reactive power change gradient
	"pgsActivePowerChangeGradient":        {Type: "U32", Gain: 1000, Addr: 42017, Quantity: 2}, // [Power grid scheduling] Active power change gradient
	"pgsScheduleInstructionValidDuration": {Type: "U32", Gain: 1, Addr: 42019, Quantity: 2},    // [Power grid scheduling] Schedule instruction valid duration
	"timeZone":                            {Type: "I16", Gain: 1, Addr: 43006, Quantity: 1},
}

// solarEdgeProps contains the complete Modbus register map for SolarEdge inverters.
// SolarEdge uses the SunSpec Modbus protocol where many values use a separate scale factor register.
// The actual value is calculated as: value = raw_value x 10^(scale_register_value)
//
// Register address ranges:
//   - 40020-40069: Device identification
//   - 40071-40107: Real-time measurements with scale factors
//
// Scale factors are stored in separate registers (ScaleAddr field) and are typically
// signed 16-bit integers representing the power of 10 to apply to the raw value.
var solarEdgeProps = map[string]defs.PropDef{
	"modelName":              {Type: "Str", Addr: 40020, Quantity: 16},
	"pn":                     {Type: "Str", Addr: 40044, Quantity: 8},
	"sn":                     {Type: "Str", Addr: 40052, Quantity: 16},
	"modelId":                {Type: "U16", Addr: 40069, Quantity: 1},
	"totalCurrent":           {Type: "U16", Addr: 40071, Quantity: 1, ScaleAddr: 40075},
	"phaseACurrent":          {Type: "U16", Addr: 40072, Quantity: 1, ScaleAddr: 40075},
	"phaseBCurrent":          {Type: "U16", Addr: 40073, Quantity: 1, ScaleAddr: 40075},
	"phaseCCurrent":          {Type: "U16", Addr: 40074, Quantity: 1, ScaleAddr: 40075},
	"LineVoltageAAndB":       {Type: "U16", Addr: 40076, Quantity: 1, ScaleAddr: 40082},
	"LineVoltageBAndC":       {Type: "U16", Addr: 40077, Quantity: 1, ScaleAddr: 40082},
	"LineVoltageCAndA":       {Type: "U16", Addr: 40078, Quantity: 1, ScaleAddr: 40082},
	"phaseAVoltage":          {Type: "U16", Addr: 40079, Quantity: 1, ScaleAddr: 40082},
	"phaseBVoltage":          {Type: "U16", Addr: 40080, Quantity: 1, ScaleAddr: 40082},
	"phaseCVoltage":          {Type: "U16", Addr: 40081, Quantity: 1, ScaleAddr: 40082},
	"activePower":            {Type: "I16", Addr: 40083, Quantity: 1, ScaleAddr: 40084},
	"gridFrequency":          {Type: "U16", Addr: 40085, Quantity: 1, ScaleAddr: 40086},
	"reactivePower":          {Type: "I16", Addr: 40089, Quantity: 1, ScaleAddr: 40090},
	"powerFactor":            {Type: "I16", Addr: 40091, Quantity: 1, ScaleAddr: 40092},
	"accumulatedEnergyYield": {Type: "U32", Addr: 40093, Quantity: 2, ScaleAddr: 40095},
	"temperature":            {Type: "I16", Addr: 40103, Quantity: 1, ScaleAddr: 40106},
	"deviceStatus": {Type: "U16", Addr: 40107, Quantity: 1, Codes: defs.Codes{
		1: "Offline",
		2: "Sleeping (auto-shutdown) - Night mode",
		3: "Grid Monitoring/wake-up",
		4: "Inverter is ON and producing power",
		5: "Production (curtailed)",
		6: "Shutting down",
		7: "Fault",
		8: "Maintenance/setup",
	}},
}
