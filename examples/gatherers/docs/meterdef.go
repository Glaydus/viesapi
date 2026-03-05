// Modbus register definitions for energy meters.
//
// # Energy Meter (NR30 Series)
//
// The NR30 series energy meters use Modbus RTU/TCP protocol with the following register map:
//
// Configuration and Identification (4147-4414):
//   - modbusPort (4147): Modbus port configuration
//   - modelNameReaded (4401): Model identifier code
//   - bootloaderVersion (4402): Bootloader firmware version (gain: 100)
//   - programVersion (4403): Program firmware version (gain: 100)
//   - nominalVoltage1 (4406): Primary nominal voltage in volts (gain: 10)
//   - nominalVoltage2 (4407): Secondary nominal voltage in volts (gain: 10)
//   - nominalCurrent1 (4408): Primary nominal current in amps (gain: 100)
//   - nominalCurrent2 (4409): Secondary nominal current in amps (gain: 100)
//   - sn76, sn54, sn32, sn10 (4411-4414): Serial number bytes
//   - sn: Complete serial number string (merged from sn76, sn54, sn32, sn10)
//
// Device Status Registers (4415-4424):
//
//   - status1 (4415): System errors and hardware status
//     Bit 0: FRAM memory corruption
//     Bit 1: no input calibration
//     Bit 4: error in configuration registers
//     Bit 5: error in displayed pages registers
//     Bit 6: error in programmable read-only register group
//     Bit 7: error of energy values
//     Bit 8: error of phase sequence
//     Bit 9: error in MQTT registries
//     Bit 10: error in supervisory relay registers
//     Bit 13: presence of Ethernet and internal memory
//     Bit 14: used battery of RTC
//
//   - status2 (4416): Alarm conditions and signaling
//     Bits 0-3: Alarm 2 condition signaling (occurrence flags)
//     Bits 4-7: Alarm 2 active conditions
//     Bits 8-11: Alarm 1 condition signaling (occurrence flags)
//     Bits 12-15: Alarm 1 active conditions
//
//   - status3 (4417): Network and archiving status
//     Bit 0: Ethernet connected
//     Bit 4: archiving in 2nd group active
//     Bit 5: archiving in 1st group active
//     Bit 7: Archiving group 2 enabled
//     Bit 8: Archiving group 1 enabled
//     Bit 9: copying to file archive from 2nd group
//     Bit 11: copying to file archive from 1st group
//     Bit 12: File archive full (less than 14 days remaining)
//     Bit 13: File archive 70% used
//     Bit 14: File archive properly initiated
//     Bit 15: Error of file archive system
//
//   - status4 (4418): Reactive power status per phase
//     Capacitive power indicators for L1, L2, L3 phases and 3-phase totals
//     Including minimum, maximum, and demand capacity flags
//
//   - status5 (4419): Alarm 1 per-phase conditions
//     Conditions 1, 2, 3 active flags for L1, L2, L3 phases
//
//   - status6 (4420): Alarm 2 per-phase conditions
//     Conditions 1, 2, 3 active flags for L1, L2, L3 phases
//
//   - status7 (4424): Protocol and relay status
//     Bit 14: MQTT protocol functions enabled
//     Bit 15: supervisory relay functions enabled
//
// Network Configuration (4421-4423):
//   - mac54, mac32, mac10: MAC address bytes
//   - mac: Complete MAC address string (merged from mac54, mac32, mac10)
//
// Energy Counters (4426-4435):
//
//	All energy values are stored as 32-bit values split across two registers (gain: 100 = 0.01 kWh resolution)
//	- activeImportedEnergy (4426-4427): Total active energy imported from grid
//	- activeExportedEnergy (4428-4429): Total active energy exported to grid
//	- reactiveInductiveEnergy (4430-4431): Total inductive reactive energy
//	- reactiveCapacitiveEnergy (4432-4433): Total capacitive reactive energy
//	- apparentEnergy (4434-4435): Total apparent energy
//
// Real-time Measurements Per Phase (7500-7523):
//
//	Each phase (L1, L2, L3) has identical register layout with 9-register spacing:
//
//	Phase L1 (7500-7505):
//	- phaseVoltageL1 (7500): Phase voltage in volts
//	- phaseCurrentL1 (7501): Phase current in amps
//	- activePowerL1 (7502): Active power in watts
//	- reactivePowerL1 (7503): Reactive power in VAR
//	- apparentPowerL1 (7504): Apparent power in VA
//	- activePowerFactorL1 (7505): Power factor
//
//	Phase L2 (7509-7514): Same parameters as L1
//	Phase L3 (7518-7523): Same parameters as L1
//
// Three-Phase Aggregated Measurements (7527-7533):
//   - avgVoltage3L (7527): Average voltage across all phases
//   - avgCurrent3L (7528): Average current across all phases
//   - activePower3L (7529): Sum of 3-phase active power (P1+P2+P3)
//   - reactivePower3L (7530): Sum of 3-phase reactive power (Q1+Q2+Q3)
//   - apparentPower3L (7531): Sum of 3-phase apparent power (S1+S2+S3)
//   - activePower3LFactorRaw (7532): 3-phase power factor (PF=P/S)
//   - avgTg3LFactor (7533): 3-phase average tangent phi factor (tg=Q/P)
//
// Grid Parameters (7536-7539):
//   - frequency (7536): Grid frequency in Hz
//   - interphaseVoltageL12 (7537): Line voltage L1-L2
//   - interphaseVoltageL23 (7538): Line voltage L2-L3
//   - interphaseVoltageL31 (7539): Line voltage L3-L1
//
// # Data Types
//
// The following data types are used in register definitions:
//   - U16: Unsigned 16-bit integer (1 register)
//   - U32: Unsigned 32-bit integer (2 registers, merged from bytes)
//   - Bytes: Raw byte data (1 register, 2 bytes)
//   - Bitfield: 16-bit bitfield with individual bit meanings
//   - Strs: Merged string array from multiple bitfield registers
//
// # Gain Values
//
// Gain represents the divisor for converting raw register values to actual values:
//   - gain=1: Direct value (no scaling)
//   - gain=10: Divide by 10 (e.g., 2350 → 235.0 V)
//   - gain=100: Divide by 100 (e.g., 12345 → 123.45 kWh)
//
// # Merged Properties
//
// Some properties combine multiple registers:
//   - sn: Complete serial number merged from four byte registers (sn76, sn54, sn32, sn10)
//   - mac: Complete MAC address merged from three byte registers (mac54, mac32, mac10)
//   - Energy counters: Merge two byte registers into U32 values
//   - status/alarms: Merge multiple bitfield registers into string arrays
//   - activePower3LFactor: Calculated from activePower3LFactorRaw
package access

import "installation/pkg/defs"

// lumelDefaultProps defines the default set of properties to read from energy meters.
// These properties provide essential operational data including identification, status,
// energy counters, and real-time power measurements.
var lumelDefaultProps = []string{
	"modelName", "deviceStatus", "modbusPort", "alarms",
	"activeImportedEnergy", "activeExportedEnergy", "reactiveInductiveEnergy",
	"reactiveCapacitiveEnergy", "apparentEnergy",
	"nominalVoltage1", "nominalVoltage2", "nominalCurrent1", "nominalCurrent2",
	"activePower3L", "reactivePower3L", "apparentPower3L", "activePower3LFactorRaw", "activePower3LFactor", "avgTg3LFactor",
}

// lumelProps contains the complete Modbus register map for NR30 series energy meters.
// Each property maps to specific Modbus registers with type information and scaling factors.
//
// Register address ranges:
//   - 4147-4424: Configuration, identification, and status
//   - 4426-4435: Energy counters (imported, exported, reactive, apparent)
//   - 7500-7523: Per-phase real-time measurements (L1, L2, L3)
//   - 7527-7533: Three-phase aggregated measurements
//   - 7536-7539: Grid parameters (frequency, interphase voltages)
//
// The meter uses byte-packed registers for serial numbers, MAC addresses, and energy counters.
// Status information is distributed across 7 bitfield registers providing detailed
// diagnostics of system health, alarms, archiving, and protocol status.
var lumelProps = map[string]defs.PropDef{
	"modbusPort":        {Addr: 4147, Quantity: 1, Type: "U16", Gain: 1},
	"modelNameReaded":   {Addr: 4401, Quantity: 1, Type: "U16", Gain: 1},
	"bootloaderVersion": {Addr: 4402, Quantity: 1, Type: "U16", Gain: 100},
	"programVersion":    {Addr: 4403, Quantity: 1, Type: "U16", Gain: 100},
	"nominalVoltage1":   {Addr: 4406, Quantity: 1, Type: "U16", Gain: 10},
	"nominalVoltage2":   {Addr: 4407, Quantity: 1, Type: "U16", Gain: 10},
	"nominalCurrent1":   {Addr: 4408, Quantity: 1, Type: "U16", Gain: 100},
	"nominalCurrent2":   {Addr: 4409, Quantity: 1, Type: "U16", Gain: 100},
	"sn76":              {Addr: 4411, Quantity: 1, Type: "Bytes", Gain: 1},
	"sn54":              {Addr: 4412, Quantity: 1, Type: "Bytes", Gain: 1},
	"sn32":              {Addr: 4413, Quantity: 1, Type: "Bytes", Gain: 1},
	"sn10":              {Addr: 4414, Quantity: 1, Type: "Bytes", Gain: 1},
	"sn": {
		Merge: []string{"sn76", "sn54", "sn32", "sn10"},
		Type:  "Str",
	},
	"status1": {Addr: 4415, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_0000_0001: {On: "FRAM memory corruption"},
			0b0000_0000_0000_0010: {On: "no input calibration"},

			0b0000_0000_0001_0000: {On: "error in configuration registers"},
			0b0000_0000_0010_0000: {On: "error in displayed pages registers"},
			0b0000_0000_0100_0000: {On: "error in configuration registers of programmable read-only register group"},
			0b0000_0000_1000_0000: {On: "error of energy values"},

			0b0000_0001_0000_0000: {On: "error of phase sequence"},
			0b0000_0010_0000_0000: {On: "error in MQTT registries"},
			0b0000_0100_0000_0000: {On: "error in the supervisory relay registers"},

			0b0010_0000_0000_0000: {On: "presence of Ethernet and internal memory"},
			0b0100_0000_0000_0000: {On: "used battery of RTC"},
		},
	},
	"status2": {Addr: 4416, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_0000_0001: {On: "signaling of condition 3 occurrence for alarm 2"},
			0b0000_0000_0000_0010: {On: "signaling of condition 2 occurrence for alarm 2"},
			0b0000_0000_0000_0100: {On: "signaling of condition 1 occurrence for alarm 2"},
			0b0000_0000_0000_1000: {On: "signaling of alarm 2 occurrence"},

			0b0000_0000_0001_0000: {On: "alarm 2 condition 3 active"},
			0b0000_0000_0010_0000: {On: "alarm 2 condition 2 active"},
			0b0000_0000_0100_0000: {On: "alarm 2 condition 1 active"},
			0b0000_0000_1000_0000: {On: "alarm 2 active"},

			0b0000_0001_0000_0000: {On: "signaling of condition 3 occurrence for alarm 1"},
			0b0000_0010_0000_0000: {On: "signaling of condition 2 occurrence for alarm 1"},
			0b0000_0100_0000_0000: {On: "signaling of condition 1 occurrence for alarm 1"},
			0b0000_1000_0000_0000: {On: "signaling of alarm 1 occurrence"},

			0b0001_0000_0000_0000: {On: "alarm 1 condition 3 active"},
			0b0010_0000_0000_0000: {On: "alarm 1 condition 2 active"},
			0b0100_0000_0000_0000: {On: "alarm 1 condition 1 active"},
			0b1000_0000_0000_0000: {On: "alarm 1 active"},
		},
	},
	"status3": {Addr: 4417, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_0000_0001: {On: "Ethernet connected"},

			0b0000_0000_0001_0000: {
				On:  "archiving in 2-nd archiving group",
				Off: "waiting until archiving conditions are met",
			},
			0b0000_0000_0010_0000: {
				On:  "archiving in 1-st archiving group",
				Off: "waiting until archiving conditions are met",
			},
			0b0000_0000_1000_0000: {On: "Archiving group 2 enabled"},

			0b0000_0001_0000_0000: {On: "Archiving group 1 enabled"},
			0b0000_0010_0000_0000: {On: "copying internal memory to file archive from 2nd archiving group"},
			0b0000_1000_0000_0000: {On: "copying internal memory to file archive from 1st archiving group"},

			0b0001_0000_0000_0000: {On: "File archive full, ( less than 14 days to complete archive filling at 1 sec. interval )"},
			0b0010_0000_0000_0000: {On: "File archive used in 70%"},
			0b0100_0000_0000_0000: {On: "File archive properly initiated"},
			0b1000_0000_0000_0000: {On: "Error of file archive system"},
		},
	},
	"status4": {Addr: 4418, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_0000_0010: {On: "Demand- capacity 3L max."},
			0b0000_0000_0000_0100: {On: "Demand- capacity 3L min."},
			0b0000_0000_0000_1000: {On: "Demand- capacity 3L"},

			0b0000_0000_0001_0000: {On: "capacitive 3L maximum"},
			0b0000_0000_0010_0000: {On: "capacitive 3L minimum"},
			0b0000_0000_0100_0000: {On: "capacitive 3L"},
			0b0000_0000_1000_0000: {On: "capacitive L3 maximum"},

			0b0000_1000_0000_0000: {On: "capacitive L3 minimum"},
			0b0000_0100_0000_0000: {On: "capacitive L3"},
			0b0000_0010_0000_0000: {On: "capacitive L2 maximum"},
			0b0000_0001_0000_0000: {On: "capacitive L2 minimum"},

			0b0001_0000_0000_0000: {On: "capacitive L2"},
			0b0010_0000_0000_0000: {On: "capacitive L1 maximum"},
			0b0100_0000_0000_0000: {On: "capacitive L1 minimum"},
			0b1000_0000_0000_0000: {On: "capacitive L1"},
		},
	},
	"status5": {Addr: 4419, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_1000_0000: {On: "alarm 1, condition 3 for L3 phase active"},

			0b0000_0001_0000_0000: {On: "alarm 1, condition 3 for L2 phase active"},
			0b0000_0010_0000_0000: {On: "alarm 1, condition 3 for L1 phase active"},
			0b0000_0100_0000_0000: {On: "alarm 1, condition 2 for L3 phase active"},
			0b0000_1000_0000_0000: {On: "alarm 1, condition 2 for L2 phase active"},

			0b0001_0000_0000_0000: {On: "alarm 1, condition 2 for L1 phase active"},
			0b0010_0000_0000_0000: {On: "alarm 1, condition 1 for L3 phase active"},
			0b0100_0000_0000_0000: {On: "alarm 1, condition 1 for L2 phase active"},
			0b1000_0000_0000_0000: {On: "alarm 1, condition 1 for L1 phase active"},
		},
	},
	"status6": {Addr: 4420, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0000_0000_1000_0000: {On: "alarm 2, condition 3 for L3 phase active"},

			0b0000_0001_0000_0000: {On: "alarm 2, condition 3 for L2 phase active"},
			0b0000_0010_0000_0000: {On: "alarm 2, condition 3 for L1 phase active"},
			0b0000_0100_0000_0000: {On: "alarm 2, condition 2 for L3 phase active"},
			0b0000_1000_0000_0000: {On: "alarm 2, condition 2 for L2 phase active"},

			0b0001_0000_0000_0000: {On: "alarm 2, condition 2 for L1 phase active"},
			0b0010_0000_0000_0000: {On: "alarm 2, condition 1 for L3 phase active"},
			0b0100_0000_0000_0000: {On: "alarm 2, condition 1 for L2 phase active"},
			0b1000_0000_0000_0000: {On: "alarm 2, condition 1 for L1 phase active"},
		},
	},
	"mac54": {Addr: 4421, Quantity: 1, Type: "Bytes", Gain: 1},
	"mac32": {Addr: 4422, Quantity: 1, Type: "Bytes", Gain: 1},
	"mac10": {Addr: 4423, Quantity: 1, Type: "Bytes", Gain: 1},
	"mac": {
		Merge: []string{"mac54", "mac32", "mac10"},
		Type:  "Str",
	},
	"status7": {Addr: 4424, Quantity: 1, Type: "Bytes", Gain: 1,
		Bits: defs.Bits{
			0b0100_0000_0000_0000: {On: "functions of MQTT protocol enabled"},
			0b1000_0000_0000_0000: {On: "functions of supervisory relay enabled"},
		},
	},
	"status": {
		Merge: []string{"status1", "status2", "status3", "status4", "status5", "status6", "status7"},
		Type:  "Strs",
	},
	"alarms": {
		Merge: []string{"status2"},
		Type:  "Strs"},
	"activeImportedEnergy32": {Addr: 4426, Quantity: 1, Type: "Bytes", Gain: 1},
	"activeImportedEnergy10": {Addr: 4427, Quantity: 1, Type: "Bytes", Gain: 1},
	"activeImportedEnergy":   {Merge: []string{"activeImportedEnergy32", "activeImportedEnergy10"}, Type: "U32", Gain: 100},

	"activeExportedEnergy32": {Addr: 4428, Quantity: 1, Type: "Bytes", Gain: 1},
	"activeExportedEnergy10": {Addr: 4429, Quantity: 1, Type: "Bytes", Gain: 1},
	"activeExportedEnergy":   {Merge: []string{"activeExportedEnergy32", "activeExportedEnergy10"}, Type: "U32", Gain: 100},

	"reactiveInductiveEnergy32": {Addr: 4430, Quantity: 1, Type: "Bytes", Gain: 1},
	"reactiveInductiveEnergy10": {Addr: 4431, Quantity: 1, Type: "Bytes", Gain: 1},
	"reactiveInductiveEnergy":   {Merge: []string{"reactiveInductiveEnergy32", "reactiveInductiveEnergy10"}, Type: "U32", Gain: 100},

	"reactiveCapacitiveEnergy32": {Addr: 4432, Quantity: 1, Type: "Bytes", Gain: 1},
	"reactiveCapacitiveEnergy10": {Addr: 4433, Quantity: 1, Type: "Bytes", Gain: 1},
	"reactiveCapacitiveEnergy":   {Merge: []string{"reactiveCapacitiveEnergy32", "reactiveCapacitiveEnergy10"}, Type: "U32", Gain: 100},

	"apparentEnergy32": {Addr: 4434, Quantity: 1, Type: "Bytes", Gain: 1},
	"apparentEnergy10": {Addr: 4435, Quantity: 1, Type: "Bytes", Gain: 1},
	"apparentEnergy":   {Merge: []string{"apparentEnergy32", "apparentEnergy10"}, Type: "U32", Gain: 100},

	"phaseVoltageL1":      {Addr: 7500, Quantity: 1, Type: "U16", Gain: 1},
	"phaseCurrentL1":      {Addr: 7501, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerL1":       {Addr: 7502, Quantity: 1, Type: "U16", Gain: 1},
	"reactivePowerL1":     {Addr: 7503, Quantity: 1, Type: "U16", Gain: 1},
	"apparentPowerL1":     {Addr: 7504, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerFactorL1": {Addr: 7505, Quantity: 1, Type: "U16", Gain: 1},

	"phaseVoltageL2":      {Addr: 7509, Quantity: 1, Type: "U16", Gain: 1},
	"phaseCurrentL2":      {Addr: 7510, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerL2":       {Addr: 7511, Quantity: 1, Type: "U16", Gain: 1},
	"reactivePowerL2":     {Addr: 7512, Quantity: 1, Type: "U16", Gain: 1},
	"apparentPowerL2":     {Addr: 7513, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerFactorL2": {Addr: 7514, Quantity: 1, Type: "U16", Gain: 1},

	"phaseVoltageL3":      {Addr: 7518, Quantity: 1, Type: "U16", Gain: 1},
	"phaseCurrentL3":      {Addr: 7519, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerL3":       {Addr: 7520, Quantity: 1, Type: "U16", Gain: 1},
	"reactivePowerL3":     {Addr: 7521, Quantity: 1, Type: "U16", Gain: 1},
	"apparentPowerL3":     {Addr: 7522, Quantity: 1, Type: "U16", Gain: 1},
	"activePowerFactorL3": {Addr: 7523, Quantity: 1, Type: "U16", Gain: 1},

	"avgVoltage3L":           {Addr: 7527, Quantity: 1, Type: "U16", Gain: 1},
	"avgCurrent3L":           {Addr: 7528, Quantity: 1, Type: "U16", Gain: 1},
	"activePower3L":          {Addr: 7529, Quantity: 1, Type: "U16", Gain: 1}, // sum of 3-phase active power (P1+P2+P3)
	"reactivePower3L":        {Addr: 7530, Quantity: 1, Type: "U16", Gain: 1}, // sum of 3-phase reactive power (Q1+Q2+Q3)
	"apparentPower3L":        {Addr: 7531, Quantity: 1, Type: "U16", Gain: 1}, // sum of 3-phase apparent power (S1+S2+S3)
	"activePower3LFactorRaw": {Addr: 7532, Quantity: 1, Type: "U16", Gain: 1}, // 3-phase active power factor (PF=P/S)
	"avgTg3LFactor":          {Addr: 7533, Quantity: 1, Type: "U16", Gain: 1}, // tg fi factor 3-phase average (tg=Q/P)

	"frequency":            {Addr: 7536, Quantity: 1, Type: "U16", Gain: 1},
	"interphaseVoltageL12": {Addr: 7537, Quantity: 1, Type: "U16", Gain: 1},
	"interphaseVoltageL23": {Addr: 7538, Quantity: 1, Type: "U16", Gain: 1},
	"interphaseVoltageL31": {Addr: 7539, Quantity: 1, Type: "U16", Gain: 1},
}
