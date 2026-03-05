# Definicje Modbus dla Falowników

Dokumentacja definicji rejestrów Modbus dla obsługiwanych modeli falowników.

## Huawei SUN2000

### Domyślne właściwości
```
modelName, sn, pn, modelId, deviceStatus, dailyEnergyYield,
accumulatedEnergyYield, activePower, reactivePower, efficiency, dayPeekActivePower,
pv1Voltage, pv1Current, pgsReactivePowerCompensationPf, pgsActivePowerDerating,
faultCode, states, alarms
```

### Rejestry Modbus

#### Informacje o urządzeniu
- **modelName** (30000-30014): String, nazwa modelu
- **sn** (30015-30024): String, numer seryjny
- **pn** (30025-30034): String, numer części
- **firmwareVer** (30035-30049): String, wersja firmware
- **softwareVer** (30050-30064): String, wersja oprogramowania
- **modelId** (30070): U16, identyfikator modelu
- **nofPvStrings** (30071): U16, liczba stringów PV
- **nofMppTrackers** (30072): U16, liczba trackerów MPP
- **ratedPower** (30073-30074): U32, moc znamionowa (gain: 1000)
- **pmax** (30075-30076): U32, maksymalna moc czynna (gain: 1000)
- **smax** (30077-30078): U32, maksymalna moc pozorna (gain: 1000)
- **qmaxFed** (30079-30080): I32, maksymalna moc bierna oddawana do sieci (gain: 1000)
- **qmaxAbsorbed** (30081-30082): I32, maksymalna moc bierna pobierana z sieci (gain: 1000)

#### Stany urządzenia (32000-32003)

**state1** (32000): Bitfield16
- Bit 0: standby
- Bit 1: grid connected
- Bit 2: grid-connected normally
- Bit 3: grid connection with derating due to power rationing
- Bit 4: grid connection with derating due to internal causes
- Bit 5: normal stop
- Bit 6: stop due to faults
- Bit 7: stop due to power rationing
- Bit 8: shutdown
- Bit 9: spot check

**state2** (32002): Bitfield16
- Bit 0: unlocked/locked
- Bit 1: PV connected/disconnected
- Bit 2: DSP data collection on/off

**state3** (32003-32004): Bitfield32
- Bit 0: off-grid/on-grid
- Bit 1: off-grid switch enabled/disabled

#### Alarmy (32008-32010)

Każdy alarm bitowy zawiera przypisany kod błędu dla szczegółowej identyfikacji usterki.

**alarm1** (32008): Bitfield16
- Bit 0: High String Input Voltage (kod: 2001)
- Bit 1: DC Arc Fault (kod: 2002)
- Bit 2: String Reverse Connection (kod: 2011)
- Bit 3: String Current Backfeed (kod: 2012)
- Bit 4: Abnormal String Power (kod: 2013)
- Bit 5: AFCI Self-Check Fail (kod: 2021)
- Bit 6: Phase Wire Short-Circuited to PE (kod: 2031)
- Bit 7: Grid Loss (kod: 2032)
- Bit 8: Grid Undervoltage (kod: 2033)
- Bit 9: Grid Overvoltage (kod: 2034)
- Bit 10: Grid Volt. Imbalance (kod: 2035)
- Bit 11: Grid Overfrequency (kod: 2036)
- Bit 12: Grid Underfrequency (kod: 2037)
- Bit 13: Unstable Grid Frequency (kod: 2038)
- Bit 14: Output Overcurrent (kod: 2039)
- Bit 15: Output DC Component Overhigh (kod: 2040)

**alarm2** (32009): Bitfield16
- Bit 0: Abnormal Residual Current (kod: 2051)
- Bit 1: Abnormal Grounding (kod: 2061)
- Bit 2: Low Insulation Resistance (kod: 2062)
- Bit 3: Overtemperature (kod: 2063)
- Bit 4: Device Fault (kod: 2064)
- Bit 5: Upgrade Failed or Version Mismatch (kod: 2065)
- Bit 6: License Expired (kod: 2066)
- Bit 7: Faulty Monitoring Unit (kod: 61440)
- Bit 8: Faulty Power Collector (kod: 2067)
- Bit 9: Battery abnormal (kod: 2068)
- Bit 10: Active Islanding (kod: 2070)
- Bit 11: Passive Islanding (kod: 2071)
- Bit 12: Transient AC Overvoltage (kod: 2072)
- Bit 13: Peripheral port short circuit (kod: 2075)
- Bit 14: Churn output overload (kod: 2077)
- Bit 15: Abnormal PV module configuration (kod: 2080)

**alarm3** (32010): Bitfield16
- Bit 0: Optimizer fault (kod: 2081)
- Bit 1: Built-in PID operation abnormal (kod: 2085)
- Bit 2: High input string voltage to ground (kod: 2014)
- Bit 3: External Fan Abnormal (kod: 2086)
- Bit 4: Battery Reverse Connection (kod: 2069)
- Bit 5: On-grid/Off-grid controller abnormal (kod: 2082)
- Bit 6: PV String Loss (kod: 2015)
- Bit 7: Internal Fan Abnormal (kod: 2087)
- Bit 8: DC Protection Unit Abnormal (kod: 2088)
- Bit 9: EL Unit Abnormal (kod: 2089)
- Bit 10: Active Adjustment Instruction Abnormal (kod: 2090)
- Bit 11: Reactive Adjustment Instruction Abnormal (kod: 2091)
- Bit 12: CT Wiring Abnormal (kod: 2092)
- Bit 13: DC Arc Fault (ADMC Alarm to be clear manually) (kod: 2003)

#### Napięcia i prądy PV (32016-32062)
Dla każdego stringa PV (1-24):
- **pvXVoltage**: I16, gain: 10
- **pvXCurrent**: I16, gain: 100

Przykład dla PV1:
- pv1Voltage (32016): I16, gain: 10
- pv1Current (32017): I16, gain: 100

#### Parametry wejściowe i sieciowe (32064-32090)
- **inputPower** (32064-32065): I32, moc wejściowa (gain: 1000)
- **gridVoltageOrLineVoltageAAndB** (32066): U16, napięcie sieci/liniowe A-B (gain: 10)
- **lineVoltageBAndC** (32067): U16, napięcie liniowe B-C (gain: 10)
- **lineVoltageCAndA** (32068): U16, napięcie liniowe C-A (gain: 10)
- **phaseAVoltage** (32069): U16, napięcie fazy A (gain: 10)
- **phaseBVoltage** (32070): U16, napięcie fazy B (gain: 10)
- **phaseCVoltage** (32071): U16, napięcie fazy C (gain: 10)
- **gridCurrentOrPhaseACurrent** (32072-32073): I32, prąd sieci/fazy A (gain: 1000)
- **phaseBCurrent** (32074-32075): I32, prąd fazy B (gain: 1000)
- **phaseCCurrent** (32076-32077): I32, prąd fazy C (gain: 1000)
- **dayPeekActivePower** (32078-32079): I32, szczytowa moc czynna dnia (gain: 1000)
- **activePower** (32080-32081): I32, moc czynna (gain: 1000)
- **reactivePower** (32082-32083): I32, moc bierna (gain: 1000)
- **powerFactor** (32084): I16, współczynnik mocy (gain: 1000)
- **gridFrequency** (32085): U16, częstotliwość sieci (gain: 100)
- **efficiency** (32086): U16, sprawność (gain: 100)
- **internalTemperature** (32087): I16, temperatura wewnętrzna (gain: 10)
- **insulationResistance** (32088): U16, rezystancja izolacji (gain: 1000)
- **deviceStatus** (32089): U16, status urządzenia (z mapowaniem Codes)
  - 0x000: Standby: initializing
  - 0x001: Standby: detecting insulation resistance
  - 0x002: Standby: detecting irradiation
  - 0x003: Standby: grid detecting
  - 0x0100: Starting
  - 0x0200: On grid
  - 0x0201: Grid connection: power limited
  - 0x0202: Grid connection: self-derating
  - 0x0203: Off-grid Running
  - 0xA000: Idle: no irradiation
  - 0xB000: Communication interrupt
  - 0xB001: Online
  - 0xC000: Uploading
- **faultCode** (32090): U16, kod błędu

#### Czasy i energia (32091-32118)
- **startupTime** (32091-32092): U32, czas startu (epoch seconds, local time)
- **shutdownTime** (32093-32094): U32, czas wyłączenia (epoch seconds, local time)
- **accumulatedEnergyYield** (32106-32107): U32, skumulowana energia (gain: 100)
- **totalInputPower** (32108-32109): U32, całkowita moc wejściowa (gain: 100)
- **dailyEnergyYield** (32114-32115): U32, dzienna energia (gain: 100)
- **monthlyEnergyYield** (32116-32117): U32, miesięczna energia (gain: 100)
- **yearlyEnergyYield** (32118-32119): U32, roczna energia (gain: 100)

#### Regulacja mocy (35300-35307)
- **activeRegulationState** (35300-35303): Bytes, stan regulacji mocy czynnej
- **reactiveRegulationState** (35304-35307): Bytes, stan regulacji mocy biernej

#### Power Meter Collection (37113)
- **pmcActivePower** (37113-37114): I32, moc czynna z licznika

#### Optymalizatory (37200-37202)
- **nofTotalOptimizers** (37200): U16, całkowita liczba optymalizatorów
- **nofOnlineOptimizers** (37201): U16, liczba optymalizatorów online
- **optimizerFeatureData** (37202): U16, dane funkcji optymalizatora
- **inverterStatus** (37518): U16, status falownika (z mapowaniem Codes)
  - 0: offline
  - 1: online

#### Konfiguracja systemu (40000+)
- **systemTime** (40000-40001): U32, czas systemowy (epoch seconds, local time)
- **pgsQuCurveMode** (40037): U16, tryb krzywej Q-U
- **pgsQuDispatchTriggerPower** (40038): U16
- **pgsFixedActivePowerDerated** (40120): U16
- **pgsReactivePowerCompensationPf** (40122): I16, kompensacja mocy biernej (PF) (gain: 1000)
- **pgsReactivePowerCompensationQs** (40123): I16, kompensacja mocy biernej (Q/S) (gain: 1000)
- **pgsActivePowerDerating** (40125): U16, obniżenie mocy czynnej (gain: 10)
- **pgsFixedActivePowerDeratedW** (40126-40127): U32, stała obniżona moc czynna (W)
- **pgsReactivePowerCompensationNightly** (40129-40130): I32, kompensacja mocy biernej nocą (gain: 1000)
- **pgsCosFiPpnCurve** (40133-40153): Mld (Bytes), krzywa cosφ-P/Pn
- **pgsQuCurve** (40154-40174): Mld (Bytes), krzywa Q-U
- **pgsPfuCurve** (40175-40195): Mld (Bytes), krzywa PF-U
- **pgsReactivePowerAdjustmentTime** (40196): U16, czas regulacji mocy biernej
- **pgsQuPowPercToExitScheduling** (40198): U16, procent mocy Q-U do wyjścia z harmonogramu
- **startup** (40200): U16, start
- **shutdown** (40201): U16, wyłączenie

#### Parametry sieci (42000+)
- **gridCode** (42000): U16, kod sieci (z mapowaniem Codes)
  - 303: Poland-EN50549-1-LV
  - 304: Poland-EN50549-1-MV
  - 305: Poland-NC-RfG-LV
- **pgsReactivePowerChangeGradient** (42015-42016): U32, gradient zmiany mocy biernej (gain: 1000)
- **pgsActivePowerChangeGradient** (42017-42018): U32, gradient zmiany mocy czynnej (gain: 1000)
- **pgsScheduleInstructionValidDuration** (42019-42020): U32, czas ważności instrukcji harmonogramu

#### Strefa czasowa (43006)
- **timeZone** (43006): I16, strefa czasowa

---

## SolarEdge

### Domyślne właściwości
```
modelName, sn, pn, modelId, deviceStatus, accumulatedEnergyYield,
activePower, reactivePower, powerFactor, temperature
```

### Rejestry Modbus

#### Informacje o urządzeniu
- **modelName** (40020-40035): String, nazwa modelu
- **pn** (40044-40051): String, numer części
- **sn** (40052-40067): String, numer seryjny
- **modelId** (40069): U16, identyfikator modelu

#### Prądy (40071-40075)
- **totalCurrent** (40071): U16, całkowity prąd (ScaleAddr: 40075)
- **phaseACurrent** (40072): U16, prąd fazy A (ScaleAddr: 40075)
- **phaseBCurrent** (40073): U16, prąd fazy B (ScaleAddr: 40075)
- **phaseCCurrent** (40074): U16, prąd fazy C (ScaleAddr: 40075)

#### Napięcia (40076-40082)
- **LineVoltageAAndB** (40076): U16, napięcie liniowe A-B (ScaleAddr: 40082)
- **LineVoltageBAndC** (40077): U16, napięcie liniowe B-C (ScaleAddr: 40082)
- **LineVoltageCAndA** (40078): U16, napięcie liniowe C-A (ScaleAddr: 40082)
- **phaseAVoltage** (40079): U16, napięcie fazy A (ScaleAddr: 40082)
- **phaseBVoltage** (40080): U16, napięcie fazy B (ScaleAddr: 40082)
- **phaseCVoltage** (40081): U16, napięcie fazy C (ScaleAddr: 40082)

#### Moc i częstotliwość (40083-40092)
- **activePower** (40083): I16, moc czynna (ScaleAddr: 40084)
- **gridFrequency** (40085): U16, częstotliwość sieci (ScaleAddr: 40086)
- **reactivePower** (40089): I16, moc bierna (ScaleAddr: 40090)
- **powerFactor** (40091): I16, współczynnik mocy (ScaleAddr: 40092)

#### Energia i temperatura (40093-40107)
- **accumulatedEnergyYield** (40093-40094): U32, skumulowana energia (ScaleAddr: 40095)
- **temperature** (40103): I16, temperatura (ScaleAddr: 40106)
- **deviceStatus** (40107): U16, status urządzenia (z mapowaniem Codes)
  - 1: Offline
  - 2: Sleeping (auto-shutdown) - Night mode
  - 3: Grid Monitoring/wake-up
  - 4: Inverter is ON and producing power
  - 5: Production (curtailed)
  - 6: Shutting down
  - 7: Fault
  - 8: Maintenance/setup

---

## Uwagi implementacyjne

### Typy danych
- **U16**: Unsigned 16-bit integer
- **I16**: Signed 16-bit integer
- **U32**: Unsigned 32-bit integer (2 rejestry)
- **I32**: Signed 32-bit integer (2 rejestry)
- **Str**: String (wielokrotność rejestrów)
- **Bitfield16**: 16-bitowe pole bitowe
- **Bitfield32**: 32-bitowe pole bitowe (2 rejestry)
- **Bytes**: Surowe bajty
- **Mld**: Alias dla Bytes (używany w starszych konfiguracjach, konwertowany automatycznie)

### Gain (współczynnik skalowania)
Wartość odczytana z rejestru musi być podzielona przez gain, aby uzyskać rzeczywistą wartość.
Przykład: jeśli gain=10 i odczytana wartość=235, rzeczywista wartość = 23.5

### Codes (mapowanie kodów)
Pole `Codes` definiuje mapowanie wartości numerycznych rejestru na czytelne dla człowieka teksty.
Używane jest dla rejestrów reprezentujących wyliczeniowe stany lub kody statusu.
Gdy wartość rejestru pasuje do klucza w mapie Codes, zwracany jest odpowiadający mu tekst.
Jeśli nie ma dopasowania, zwracany jest sformatowany kod (np. "Code: 0x123").

Przykład użycia w rejestrze deviceStatus (32089):
- Wartość 0x0200 zwraca "On grid"
- Wartość 0x0201 zwraca "Grid connection: power limited"
- Wartość 0xFFFF (brak w mapie) zwraca "Code: 0xFFFF"

### ScaleAddr (SolarEdge)
Niektóre rejestry SolarEdge używają osobnego rejestru do przechowywania wykładnika potęgi 10.
Wartość = odczytana_wartość × 10^(wartość_z_ScaleAddr)

### Merged Properties
Właściwości "states" i "alarms" są złożeniem wielu pól bitowych w jedną tablicę stringów.

### Bits z kodami błędów
Dla rejestrów alarmowych (alarm1, alarm2, alarm3), każdy bit zawiera nie tylko opis alarmu (pole `on`),
ale również przypisany kod błędu (pole `code`). Kody te są zwracane jako część struktury Alarm
zawierającej nazwę alarmu i jego kod numeryczny.

---

*Szczegóły implementacji: inverterdef.go*
