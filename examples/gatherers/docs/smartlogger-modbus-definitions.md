# Definicje Modbus dla Huawei SmartLogger

Dokumentacja definicji rejestrów Modbus dla urządzeń Huawei SmartLogger 3000 oraz podłączonych do nich falowników.

## SmartLogger 3000

SmartLogger 3000 to urządzenie monitorujące i zarządzające instalacjami fotowoltaicznymi Huawei.
Komunikuje się z falownikami i innymi urządzeniami, agregując dane i udostępniając je przez protokół Modbus TCP.

### Domyślne właściwości
```
deviceStatus, accumulatedEnergyYield, dailyEnergyYield, dailyPowerDuration,
activePower, inputPower, reactivePower, co2Reduction
```

### Rejestry Modbus

#### Czas systemowy (40000-40009)
- **dateTime** (40000-40001): U32, data i czas (epoch seconds)
- **localTime** (40009-40010): U32, czas lokalny (epoch seconds)

#### Pomiary mocy (40521-40544)
- **inputPower** (40521-40522): U32, moc wejściowa DC (gain: 1000)
- **co2Reduction** (40523-40524): U32, redukcja CO2 w kg (gain: 10)
- **activePower** (40525-40526): I32, moc czynna AC (gain: 1000)
- **powerFactor** (40532): I16, współczynnik mocy (gain: 1000)
- **reactivePower** (40544-40545): I32, moc bierna (gain: 1000)

#### Liczniki energii (40560-40564)
- **accumulatedEnergyYield** (40560-40561): U32, skumulowana energia w kWh (gain: 10)
- **dailyEnergyYield** (40562-40563): U32, dzienna energia w kWh (gain: 10)
- **dailyPowerDuration** (40564-40565): U32, dzienny czas pracy w minutach (gain: 10)

#### Prądy fazowe (40572-40574)
- **phaseACurrent** (40572): I16, prąd fazy A w amperach
- **phaseBCurrent** (40573): I16, prąd fazy B w amperach
- **phaseCCurrent** (40574): I16, prąd fazy C w amperach

#### Identyfikacja (40713-40736)
- **esn** (40713-40722): String, numer seryjny urządzenia
- **deviceAccessStatus** (40736): U16, status dostępu do urządzenia

#### Alarmy (50000-50001)

Każdy alarm bitowy zawiera przypisany kod błędu dla szczegółowej identyfikacji usterki.

**alarm1** (50000): Bitfield16
- Bit 3: Abnormal Active Schedule (kod: 1100)
- Bit 11: Abnormal Reactive Schedule (kod: 1101)

**alarm2** (50001): Bitfield16
- Bit 0: MCB Disconnect (kod: 1103)
- Bit 1: Abnormal Cubicle (kod: 1104)
- Bit 2: Device Address Conflict (kod: 1105)
- Bit 3: AC SPD fault (kod: 1106)
- Bit 13: 24V power failure (kod: 1115)
- Bit 14: License Expired (kod: 1119)

#### Status urządzenia (65534)
- **deviceStatus** (65534): U16, status urządzenia (z mapowaniem Codes)
  - 0xB000: Disconnection
  - 0xB001: Online

---

## SmartLogger - Falowniki

Rejestry falowników podłączonych do SmartLoggera są dostępne przez adresację względną.
Każdy falownik ma swój własny zestaw rejestrów zaczynający się od adresu 0.

### Rejestry Modbus falownika

#### Pomiary mocy (0-8)
- **activePower** (0-1): I32, moc czynna w W (gain: 1000)
- **reactivePower** (2-3): I32, moc bierna w VAr (gain: 1000)
- **inputDc** (4): I16, prąd DC wejściowy w A (gain: 100)
- **inputPower** (5-6): U32, moc wejściowa DC w W (gain: 1000)
- **insulationResistance** (7): U16, rezystancja izolacji w kΩ (gain: 1000)
- **powerFactor** (8): I16, współczynnik mocy (gain: 1000)

#### Status i temperatura (9-11)
- **deviceStatus** (9): U16, status urządzenia (z mapowaniem Codes)
  - 0xB000: Communication interrupt
  - 0xB001: Online
  - 0xC000: Uploading
- **temperature** (11): I16, temperatura w °C (gain: 10)

#### Kody błędów (12-18)
- **majorFaultCode** (12-13): U32, kod głównego błędu
- **minorFaultCode** (14-15): U32, kod drugorzędnego błędu
- **warningCode** (16-17): U32, kod ostrzeżenia

---

## Uwagi implementacyjne

### Typy danych
- **U16**: Unsigned 16-bit integer
- **I16**: Signed 16-bit integer
- **U32**: Unsigned 32-bit integer (2 rejestry)
- **I32**: Signed 32-bit integer (2 rejestry)
- **Str**: String (wielokrotność rejestrów)
- **Bitfield16**: 16-bitowe pole bitowe
- **Strs**: Tablica stringów (dla merged properties)

### Gain (współczynnik skalowania)
Wartość odczytana z rejestru musi być podzielona przez gain, aby uzyskać rzeczywistą wartość.
Przykład: jeśli gain=10 i odczytana wartość=235, rzeczywista wartość = 23.5

### Codes (mapowanie kodów)
Pole `Codes` definiuje mapowanie wartości numerycznych rejestru na czytelne dla człowieka teksty.
Używane jest dla rejestrów reprezentujących wyliczeniowe stany lub kody statusu.
Gdy wartość rejestru pasuje do klucza w mapie Codes, zwracany jest odpowiadający mu tekst.
Jeśli nie ma dopasowania, zwracany jest sformatowany kod (np. "Code: 0x123").

Przykład użycia w rejestrze deviceStatus (65534):
- Wartość 0xB000 zwraca "Disconnection"
- Wartość 0xB001 zwraca "Online"
- Wartość 0xFFFF (brak w mapie) zwraca "Code: 0xFFFF"

### Merged Properties
Właściwość "alarms" jest złożeniem wielu pól bitowych (alarm1, alarm2) w jedną tablicę stringów.

### Bits z kodami błędów
Dla rejestrów alarmowych (alarm1, alarm2), każdy bit zawiera nie tylko opis alarmu (pole `on`),
ale również przypisany kod błędu (pole `code`). Kody te są zwracane jako część struktury Alarm
zawierającej nazwę alarmu i jego kod numeryczny.

### Adresacja falowników
SmartLogger zarządza wieloma falownikami. Dostęp do rejestrów konkretnego falownika wymaga
odpowiedniej adresacji przez protokół Modbus (zazwyczaj przez slave ID lub offset adresowy).

---

*Szczegóły implementacji: smartdef.go*
