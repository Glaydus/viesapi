# Definicje Modbus dla Liczników Energii

Dokumentacja definicji rejestrów Modbus dla liczników energii serii Lumel NR30.

## NR30 Series Energy Meter

### Domyślne właściwości
```
modelName, deviceStatus, modbusPort, alarms,
activeImportedEnergy, activeExportedEnergy, reactiveInductiveEnergy,
reactiveCapacitiveEnergy, apparentEnergy,
nominalVoltage1, nominalVoltage2, nominalCurrent1, nominalCurrent2,
activePower3L, reactivePower3L, apparentPower3L, activePower3LFactorRaw,
activePower3LFactor, avgTg3LFactor
```

### Rejestry Modbus

#### Konfiguracja i identyfikacja (4147-4414)

- **modbusPort** (4147): U16, konfiguracja portu Modbus
- **modelNameReaded** (4401): U16, kod identyfikacyjny modelu
- **bootloaderVersion** (4402): U16, wersja bootloadera (gain: 100)
- **programVersion** (4403): U16, wersja programu (gain: 100)
- **nominalVoltage1** (4406): U16, napięcie znamionowe pierwotne w V (gain: 10)
- **nominalVoltage2** (4407): U16, napięcie znamionowe wtórne w V (gain: 10)
- **nominalCurrent1** (4408): U16, prąd znamionowy pierwotny w A (gain: 100)
- **nominalCurrent2** (4409): U16, prąd znamionowy wtórny w A (gain: 100)
- **sn76, sn54, sn32, sn10** (4411-4414): Bytes, bajty numeru seryjnego
- **sn**: String, pełny numer seryjny (scalony z sn76, sn54, sn32, sn10)

#### Rejestry statusu urządzenia (4415-4424)

**status1** (4415): Błędy systemowe i status sprzętu
- Bit 0: Uszkodzenie pamięci FRAM
- Bit 1: Brak kalibracji wejścia
- Bit 4: Błąd w rejestrach konfiguracyjnych
- Bit 5: Błąd w rejestrach wyświetlanych stron
- Bit 6: Błąd w rejestrach grupy tylko do odczytu
- Bit 7: Błąd wartości energii
- Bit 8: Błąd sekwencji faz
- Bit 9: Błąd w rejestrach MQTT
- Bit 10: Błąd w rejestrach przekaźnika nadzorującego
- Bit 13: Obecność Ethernet i pamięci wewnętrznej
- Bit 14: Używana bateria RTC

**status2** (4416): Warunki alarmowe i sygnalizacja
- Bity 0-3: Sygnalizacja alarmu 2 (flagi wystąpienia)
  - Bit 0: Sygnalizacja wystąpienia warunku 3 dla alarmu 2
  - Bit 1: Sygnalizacja wystąpienia warunku 2 dla alarmu 2
  - Bit 2: Sygnalizacja wystąpienia warunku 1 dla alarmu 2
  - Bit 3: Sygnalizacja wystąpienia alarmu 2
- Bity 4-7: Aktywne warunki alarmu 2
  - Bit 4: Warunek 3 alarmu 2 aktywny
  - Bit 5: Warunek 2 alarmu 2 aktywny
  - Bit 6: Warunek 1 alarmu 2 aktywny
  - Bit 7: Alarm 2 aktywny
- Bity 8-11: Sygnalizacja alarmu 1 (flagi wystąpienia)
  - Bit 8: Sygnalizacja wystąpienia warunku 3 dla alarmu 1
  - Bit 9: Sygnalizacja wystąpienia warunku 2 dla alarmu 1
  - Bit 10: Sygnalizacja wystąpienia warunku 1 dla alarmu 1
  - Bit 11: Sygnalizacja wystąpienia alarmu 1
- Bity 12-15: Aktywne warunki alarmu 1
  - Bit 12: Warunek 3 alarmu 1 aktywny
  - Bit 13: Warunek 2 alarmu 1 aktywny
  - Bit 14: Warunek 1 alarmu 1 aktywny
  - Bit 15: Alarm 1 aktywny

**status3** (4417): Status sieci i archiwizacji
- Bit 0: Ethernet podłączony
- Bit 4: Archiwizacja w 2. grupie aktywna
- Bit 5: Archiwizacja w 1. grupie aktywna
- Bit 7: Grupa archiwizacji 2 włączona
- Bit 8: Grupa archiwizacji 1 włączona
- Bit 9: Kopiowanie do archiwum plików z 2. grupy
- Bit 11: Kopiowanie do archiwum plików z 1. grupy
- Bit 12: Archiwum plików pełne (mniej niż 14 dni do zapełnienia przy interwale 1s)
- Bit 13: Archiwum plików wykorzystane w 70%
- Bit 14: Archiwum plików prawidłowo zainicjowane
- Bit 15: Błąd systemu archiwum plików

**status4** (4418): Status mocy biernej na fazę
- Wskaźniki mocy pojemnościowej dla faz L1, L2, L3 i sum 3-fazowych
- Zawiera flagi minimum, maksimum i zapotrzebowania

**status5** (4419): Warunki alarmu 1 na fazę
- Bit 7: Alarm 1, warunek 3 dla fazy L3 aktywny
- Bit 8: Alarm 1, warunek 3 dla fazy L2 aktywny
- Bit 9: Alarm 1, warunek 3 dla fazy L1 aktywny
- Bit 10: Alarm 1, warunek 2 dla fazy L3 aktywny
- Bit 11: Alarm 1, warunek 2 dla fazy L2 aktywny
- Bit 12: Alarm 1, warunek 2 dla fazy L1 aktywny
- Bit 13: Alarm 1, warunek 1 dla fazy L3 aktywny
- Bit 14: Alarm 1, warunek 1 dla fazy L2 aktywny
- Bit 15: Alarm 1, warunek 1 dla fazy L1 aktywny

**status6** (4420): Warunki alarmu 2 na fazę
- Bit 7: Alarm 2, warunek 3 dla fazy L3 aktywny
- Bit 8: Alarm 2, warunek 3 dla fazy L2 aktywny
- Bit 9: Alarm 2, warunek 3 dla fazy L1 aktywny
- Bit 10: Alarm 2, warunek 2 dla fazy L3 aktywny
- Bit 11: Alarm 2, warunek 2 dla fazy L2 aktywny
- Bit 12: Alarm 2, warunek 2 dla fazy L1 aktywny
- Bit 13: Alarm 2, warunek 1 dla fazy L3 aktywny
- Bit 14: Alarm 2, warunek 1 dla fazy L2 aktywny
- Bit 15: Alarm 2, warunek 1 dla fazy L1 aktywny

**status7** (4424): Status protokołu i przekaźnika
- Bit 14: Funkcje protokołu MQTT włączone
- Bit 15: Funkcje przekaźnika nadzorującego włączone

#### Konfiguracja sieciowa (4421-4423)
- **mac54, mac32, mac10** (4421-4423): Bytes, bajty adresu MAC
- **mac**: String, pełny adres MAC (scalony z mac54, mac32, mac10)

#### Liczniki energii (4426-4435)

Wszystkie wartości energii są przechowywane jako 32-bitowe wartości podzielone na dwa rejestry (gain: 100 = rozdzielczość 0.01 kWh)

- **activeImportedEnergy** (4426-4427): U32, całkowita energia czynna pobrana z sieci (gain: 100)
- **activeExportedEnergy** (4428-4429): U32, całkowita energia czynna oddana do sieci (gain: 100)
- **reactiveInductiveEnergy** (4430-4431): U32, całkowita energia bierna indukcyjna (gain: 100)
- **reactiveCapacitiveEnergy** (4432-4433): U32, całkowita energia bierna pojemnościowa (gain: 100)
- **apparentEnergy** (4434-4435): U32, całkowita energia pozorna (gain: 100)

#### Pomiary w czasie rzeczywistym na fazę (7500-7523)

Każda faza (L1, L2, L3) ma identyczny układ rejestrów z odstępem 9 rejestrów:

**Faza L1** (7500-7505):
- **phaseVoltageL1** (7500): U16, napięcie fazy w V
- **phaseCurrentL1** (7501): U16, prąd fazy w A
- **activePowerL1** (7502): U16, moc czynna w W
- **reactivePowerL1** (7503): U16, moc bierna w VAR
- **apparentPowerL1** (7504): U16, moc pozorna w VA
- **activePowerFactorL1** (7505): U16, współczynnik mocy

**Faza L2** (7509-7514): Te same parametry co L1

**Faza L3** (7518-7523): Te same parametry co L1

#### Pomiary zagregowane 3-fazowe (7527-7533)

- **avgVoltage3L** (7527): U16, średnie napięcie wszystkich faz
- **avgCurrent3L** (7528): U16, średni prąd wszystkich faz
- **activePower3L** (7529): U16, suma mocy czynnej 3-fazowej (P1+P2+P3)
- **reactivePower3L** (7530): U16, suma mocy biernej 3-fazowej (Q1+Q2+Q3)
- **apparentPower3L** (7531): U16, suma mocy pozornej 3-fazowej (S1+S2+S3)
- **activePower3LFactorRaw** (7532): U16, współczynnik mocy 3-fazowej (PF=P/S)
- **avgTg3LFactor** (7533): U16, średni współczynnik tangens phi 3-fazowy (tg=Q/P)

#### Parametry sieci (7536-7539)

- **frequency** (7536): U16, częstotliwość sieci w Hz
- **interphaseVoltageL12** (7537): U16, napięcie międzyfazowe L1-L2
- **interphaseVoltageL23** (7538): U16, napięcie międzyfazowe L2-L3
- **interphaseVoltageL31** (7539): U16, napięcie międzyfazowe L3-L1

---

## Uwagi implementacyjne

### Typy danych
- **U16**: Unsigned 16-bit integer (1 rejestr)
- **U32**: Unsigned 32-bit integer (2 rejestry, scalane z bajtów)
- **Bytes**: Surowe dane bajtowe (1 rejestr, 2 bajty)
- **Bitfield**: 16-bitowe pole bitowe z indywidualnymi znaczeniami bitów
- **Strs**: Scalona tablica stringów z wielu rejestrów bitowych

### Gain (współczynnik skalowania)
Wartość odczytana z rejestru musi być podzielona przez gain, aby uzyskać rzeczywistą wartość:
- gain=1: Wartość bezpośrednia (bez skalowania)
- gain=10: Dzielenie przez 10 (np. 2350 → 235.0 V)
- gain=100: Dzielenie przez 100 (np. 12345 → 123.45 kWh)

### Właściwości scalone (Merged Properties)
Niektóre właściwości łączą wiele rejestrów:
- **sn**: Scalony numer seryjny z czterech rejestrów bajtowych (sn76, sn54, sn32, sn10) w string
- **mac**: Scalony adres MAC z trzech rejestrów bajtowych (mac54, mac32, mac10) w string
- Liczniki energii: Scalają dwa rejestry bajtowe w wartości U32
- status/alarms: Scalają wiele rejestrów bitowych w tablice stringów
- activePower3LFactor: Obliczany z activePower3LFactorRaw

### Układ rejestrów fazowych
Pomiary dla każdej fazy (L1, L2, L3) są rozmieszczone w regularnych odstępach:
- L1: 7500-7505
- L2: 7509-7514 (offset +9)
- L3: 7518-7523 (offset +18)

---

*Sczegóły implementacji: meterdef.go*
