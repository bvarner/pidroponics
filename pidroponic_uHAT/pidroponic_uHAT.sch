EESchema Schematic File Version 4
LIBS:pidroponic_uHAT-cache
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "Pidroponics uHAT"
Date "2020-04-15"
Rev "1.0"
Comp ""
Comment1 "This Schematic is licensed under Apache License Version 2.0"
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Connector_Generic:Conn_02x20_Odd_Even J1
U 1 1 5C77771F
P 2250 2050
F 0 "J1" H 2300 3167 50  0000 C CNN
F 1 "GPIO_CONNECTOR" H 2300 3076 50  0000 C CNN
F 2 "lib:PinSocket_2x20_P2.54mm_Vertical_Centered_Anchor" H 2250 2050 50  0001 C CNN
F 3 "~" H 2250 2050 50  0001 C CNN
	1    2250 2050
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0101
U 1 1 5C777805
P 1850 3200
F 0 "#PWR0101" H 1850 2950 50  0001 C CNN
F 1 "GND" H 1855 3027 50  0001 C CNN
F 2 "" H 1850 3200 50  0001 C CNN
F 3 "" H 1850 3200 50  0001 C CNN
	1    1850 3200
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0102
U 1 1 5C777838
P 2750 3200
F 0 "#PWR0102" H 2750 2950 50  0001 C CNN
F 1 "GND" H 2755 3027 50  0001 C CNN
F 2 "" H 2750 3200 50  0001 C CNN
F 3 "" H 2750 3200 50  0001 C CNN
	1    2750 3200
	1    0    0    -1  
$EndComp
Wire Wire Line
	2050 1550 1850 1550
Wire Wire Line
	2550 1350 2750 1350
Wire Wire Line
	2750 1350 2750 1750
Wire Wire Line
	2550 1750 2750 1750
Connection ~ 2750 1750
Wire Wire Line
	2750 1750 2750 2050
Wire Wire Line
	2550 2050 2750 2050
Connection ~ 2750 2050
Wire Wire Line
	2550 2550 2750 2550
Wire Wire Line
	2750 2050 2750 2550
Connection ~ 2750 2550
Wire Wire Line
	2750 2550 2750 2750
Wire Wire Line
	2550 2750 2750 2750
Connection ~ 2750 2750
Wire Wire Line
	2750 2750 2750 3200
$Comp
L power:+3.3V #PWR0103
U 1 1 5C777AB0
P 1800 1050
F 0 "#PWR0103" H 1800 900 50  0001 C CNN
F 1 "+3.3V" H 1815 1223 50  0000 C CNN
F 2 "" H 1800 1050 50  0001 C CNN
F 3 "" H 1800 1050 50  0001 C CNN
	1    1800 1050
	1    0    0    -1  
$EndComp
Wire Wire Line
	1800 1150 1800 1050
Wire Wire Line
	2050 1950 1800 1950
Wire Wire Line
	1800 1950 1800 1150
Connection ~ 1800 1150
$Comp
L power:+5V #PWR0104
U 1 1 5C777E01
P 2850 1050
F 0 "#PWR0104" H 2850 900 50  0001 C CNN
F 1 "+5V" H 2865 1223 50  0000 C CNN
F 2 "" H 2850 1050 50  0001 C CNN
F 3 "" H 2850 1050 50  0001 C CNN
	1    2850 1050
	1    0    0    -1  
$EndComp
Wire Wire Line
	2550 1150 2850 1150
Wire Wire Line
	2850 1150 2850 1050
Wire Wire Line
	2550 1250 2850 1250
Wire Wire Line
	2850 1250 2850 1150
Connection ~ 2850 1150
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5C77824A
P 1400 1050
F 0 "#FLG0101" H 1400 1125 50  0001 C CNN
F 1 "PWR_FLAG" H 1400 1224 50  0000 C CNN
F 2 "" H 1400 1050 50  0001 C CNN
F 3 "~" H 1400 1050 50  0001 C CNN
	1    1400 1050
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0105
U 1 1 5C778504
P 1450 3300
F 0 "#PWR0105" H 1450 3050 50  0001 C CNN
F 1 "GND" H 1455 3127 50  0001 C CNN
F 2 "" H 1450 3300 50  0001 C CNN
F 3 "" H 1450 3300 50  0001 C CNN
	1    1450 3300
	1    0    0    -1  
$EndComp
Wire Wire Line
	1450 3250 1450 3300
Wire Wire Line
	1800 1150 2050 1150
Wire Wire Line
	1400 1150 1400 1050
Wire Wire Line
	1400 1150 1800 1150
$Comp
L power:PWR_FLAG #FLG0103
U 1 1 5C77CEFA
P 3200 1050
F 0 "#FLG0103" H 3200 1125 50  0001 C CNN
F 1 "PWR_FLAG" H 3200 1224 50  0000 C CNN
F 2 "" H 3200 1050 50  0001 C CNN
F 3 "~" H 3200 1050 50  0001 C CNN
	1    3200 1050
	1    0    0    -1  
$EndComp
Wire Wire Line
	2850 1150 3200 1150
Wire Wire Line
	3200 1050 3200 1150
Text Label 1100 1250 0    50   ~ 0
GPIO2_SDA1
Text Label 1100 1350 0    50   ~ 0
GPIO3_SCL1
Text Label 1100 1450 0    50   ~ 0
GPIO4_GPIO_GCLK
Text Label 1100 1750 0    50   ~ 0
GPIO27_GEN2
Text Label 1100 1850 0    50   ~ 0
GPIO22_GEN3
Text Label 1100 2050 0    50   ~ 0
GPIO10_SPI_MOSI
Wire Wire Line
	1000 2050 2050 2050
Wire Wire Line
	1000 2150 2050 2150
Wire Wire Line
	1000 2250 2050 2250
Wire Wire Line
	1000 2450 2050 2450
Wire Wire Line
	1000 2550 2050 2550
Wire Wire Line
	1000 2650 2050 2650
Wire Wire Line
	1000 2750 2050 2750
Wire Wire Line
	1000 2850 2050 2850
Wire Wire Line
	1000 2950 2050 2950
Wire Wire Line
	1000 1850 2050 1850
Wire Wire Line
	1000 1750 2050 1750
Wire Wire Line
	1000 1650 2050 1650
Wire Wire Line
	1000 1450 2050 1450
Wire Wire Line
	1000 1350 2050 1350
Wire Wire Line
	1000 1250 2050 1250
Text Label 1100 2150 0    50   ~ 0
GPIO9_SPI_MISO
Text Label 1100 2250 0    50   ~ 0
GPIO11_SPI_SCLK
Text Label 1000 2550 0    50   ~ 0
GPIO5
Text Label 1000 2650 0    50   ~ 0
GPIO6
Text Label 1000 2750 0    50   ~ 0
GPIO13
Text Label 1000 2850 0    50   ~ 0
GPIO19
Text Label 1000 2950 0    50   ~ 0
GPIO26
NoConn ~ 1000 1750
NoConn ~ 1000 1850
NoConn ~ 1000 2050
NoConn ~ 1000 2150
NoConn ~ 1000 2250
Text Label 2900 1450 0    50   ~ 0
GPIO14_TXD0
Text Label 2900 1550 0    50   ~ 0
GPIO15_RXD0
Text Label 2900 1650 0    50   ~ 0
GPIO18_GEN1
Text Label 2900 1850 0    50   ~ 0
GPIO23_GEN4
Text Label 2900 1950 0    50   ~ 0
GPIO24_GEN5
Text Label 3600 2150 0    50   ~ 0
GPIO25
Text Label 2900 2250 0    50   ~ 0
GPIO8_SPI_CE0_N
Text Label 2900 2350 0    50   ~ 0
GPIO7_SPI_CE1_N
Wire Wire Line
	2550 2250 3600 2250
Wire Wire Line
	2550 2350 3600 2350
Text Label 3600 2650 0    50   ~ 0
GPIO12
Text Label 3600 2850 0    50   ~ 0
GPIO16
Text Label 3600 2950 0    50   ~ 0
GPIO20
Text Label 3600 3050 0    50   ~ 0
GPIO21
Wire Wire Line
	2550 1450 3600 1450
Wire Wire Line
	2550 1550 3600 1550
Wire Wire Line
	2550 1650 3600 1650
Wire Wire Line
	2550 1850 3600 1850
Wire Wire Line
	2550 1950 3600 1950
Wire Wire Line
	2550 2150 3600 2150
Wire Wire Line
	2550 2450 3600 2450
Wire Wire Line
	2550 2650 3600 2650
Wire Wire Line
	2550 2850 3600 2850
Wire Wire Line
	2550 2950 3600 2950
NoConn ~ 3600 1850
NoConn ~ 3600 1950
NoConn ~ 3600 2250
NoConn ~ 3600 2350
Wire Wire Line
	2550 3050 3600 3050
$Comp
L Mechanical:MountingHole H1
U 1 1 5C7C4C81
P 900 6950
F 0 "H1" H 1000 6996 50  0000 L CNN
F 1 "MountingHole" H 1000 6905 50  0000 L CNN
F 2 "lib:MountingHole_2.7mm_M2.5_uHAT_RPi" H 900 6950 50  0001 C CNN
F 3 "~" H 900 6950 50  0001 C CNN
	1    900  6950
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H2
U 1 1 5C7C7FBC
P 900 7150
F 0 "H2" H 1000 7196 50  0000 L CNN
F 1 "MountingHole" H 1000 7105 50  0000 L CNN
F 2 "lib:MountingHole_2.7mm_M2.5_uHAT_RPi" H 900 7150 50  0001 C CNN
F 3 "~" H 900 7150 50  0001 C CNN
	1    900  7150
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H3
U 1 1 5C7C8014
P 900 7350
F 0 "H3" H 1000 7396 50  0000 L CNN
F 1 "MountingHole" H 1000 7305 50  0000 L CNN
F 2 "lib:MountingHole_2.7mm_M2.5_uHAT_RPi" H 900 7350 50  0001 C CNN
F 3 "~" H 900 7350 50  0001 C CNN
	1    900  7350
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H4
U 1 1 5C7C8030
P 900 7550
F 0 "H4" H 1000 7596 50  0000 L CNN
F 1 "MountingHole" H 1000 7505 50  0000 L CNN
F 2 "lib:MountingHole_2.7mm_M2.5_uHAT_RPi" H 900 7550 50  0001 C CNN
F 3 "~" H 900 7550 50  0001 C CNN
	1    900  7550
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 5E9778F9
P 1600 6000
F 0 "C1" H 1715 6046 50  0000 L CNN
F 1 "10uf" H 1715 5955 50  0000 L CNN
F 2 "Capacitor_SMD:C_0805_2012Metric" H 1638 5850 50  0001 C CNN
F 3 "~" H 1600 6000 50  0001 C CNN
F 4 "CAP CER 10UF 16V X5R 0805" H 1600 6000 50  0001 C CNN "Description"
F 5 "445-7644-1-ND" H 1600 6000 50  0001 C CNN "Digi-Key_PN"
	1    1600 6000
	1    0    0    -1  
$EndComp
$Comp
L Regulator_Linear:LM7805_TO220 U1
U 1 1 5E97813E
P 2150 5850
F 0 "U1" H 2150 6092 50  0000 C CNN
F 1 "LM7805_TO220" H 2150 6001 50  0000 C CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2150 6075 50  0001 C CIN
F 3 "http://www.fairchildsemi.com/ds/LM/LM7805.pdf" H 2150 5800 50  0001 C CNN
F 4 "IC REG LINEAR 5V 1A TO220-3" H 2150 5850 50  0001 C CNN "Description"
F 5 "296-47192-ND" H 2150 5850 50  0001 C CNN "Digi-Key_PN"
	1    2150 5850
	1    0    0    -1  
$EndComp
$Comp
L Device:C C2
U 1 1 5E979431
P 2650 6000
F 0 "C2" H 2765 6046 50  0000 L CNN
F 1 "10uf" H 2765 5955 50  0000 L CNN
F 2 "Capacitor_SMD:C_0805_2012Metric" H 2688 5850 50  0001 C CNN
F 3 "~" H 2650 6000 50  0001 C CNN
F 4 "CAP CER 10UF 16V X5R 0805" H 2650 6000 50  0001 C CNN "Description"
F 5 "445-7644-1-ND" H 2650 6000 50  0001 C CNN "Digi-Key_PN"
	1    2650 6000
	1    0    0    -1  
$EndComp
$Comp
L dk_Transistors-FETs-MOSFETs-Single:DMG2305UX-13 Q1
U 1 1 5E980B0B
P 3750 5850
F 0 "Q1" V 4017 5850 60  0000 C CNN
F 1 "DMG2305UX-13" V 3911 5850 60  0000 C CNN
F 2 "Package_TO_SOT_SMD:SOT-23" H 3950 6050 60  0001 L CNN
F 3 "https://www.diodes.com/assets/Datasheets/DMG2305UX.pdf" H 3950 6150 60  0001 L CNN
F 4 "DMG2305UX-13DICT-ND" H 3950 6250 60  0001 L CNN "Digi-Key_PN"
F 5 "DMG2305UX-13" H 3950 6350 60  0001 L CNN "MPN"
F 6 "Discrete Semiconductor Products" H 3950 6450 60  0001 L CNN "Category"
F 7 "Transistors - FETs, MOSFETs - Single" H 3950 6550 60  0001 L CNN "Family"
F 8 "https://www.diodes.com/assets/Datasheets/DMG2305UX.pdf" H 3950 6650 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/diodes-incorporated/DMG2305UX-13/DMG2305UX-13DICT-ND/4251589" H 3950 6750 60  0001 L CNN "DK_Detail_Page"
F 10 "MOSFET P-CH 20V 4.2A SOT23" H 3950 6850 60  0001 L CNN "Description"
F 11 "Diodes Incorporated" H 3950 6950 60  0001 L CNN "Manufacturer"
F 12 "Active" H 3950 7050 60  0001 L CNN "Status"
	1    3750 5850
	0    -1   -1   0   
$EndComp
$Comp
L Device:R R1
U 1 1 5E983568
P 3500 7150
F 0 "R1" H 3570 7196 50  0000 L CNN
F 1 "10K" H 3570 7105 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 3430 7150 50  0001 C CNN
F 3 "~" H 3500 7150 50  0001 C CNN
F 4 "RES 10K OHM 0.5% 1/8W 0805" H 3500 7150 50  0001 C CNN "Description"
F 5 "RNCF0805DTE10K0CT-ND" H 3500 7150 50  0001 C CNN "Digi-Key_PN"
	1    3500 7150
	1    0    0    -1  
$EndComp
$Comp
L Device:R R2
U 1 1 5E983A86
P 4200 7150
F 0 "R2" H 4270 7196 50  0000 L CNN
F 1 "47K" H 4270 7105 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 4130 7150 50  0001 C CNN
F 3 "~" H 4200 7150 50  0001 C CNN
F 4 "RES 47K OHM 0.5% 1/8W 0805" H 4200 7150 50  0001 C CNN "Description"
F 5 "RNCF0805DTE47K0CT-ND" H 4200 7150 50  0001 C CNN "Digi-Key_PN"
	1    4200 7150
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0106
U 1 1 5E9840F3
P 3500 7400
F 0 "#PWR0106" H 3500 7150 50  0001 C CNN
F 1 "GND" H 3505 7227 50  0000 C CNN
F 2 "" H 3500 7400 50  0001 C CNN
F 3 "" H 3500 7400 50  0001 C CNN
	1    3500 7400
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0107
U 1 1 5E984533
P 4200 7400
F 0 "#PWR0107" H 4200 7150 50  0001 C CNN
F 1 "GND" H 4205 7227 50  0000 C CNN
F 2 "" H 4200 7400 50  0001 C CNN
F 3 "" H 4200 7400 50  0001 C CNN
	1    4200 7400
	1    0    0    -1  
$EndComp
Wire Wire Line
	1050 5950 1050 6150
$Comp
L power:PWR_FLAG #FLG0104
U 1 1 5E988141
P 1600 5850
F 0 "#FLG0104" H 1600 5925 50  0001 C CNN
F 1 "PWR_FLAG" H 1600 6023 50  0000 C CNN
F 2 "" H 1600 5850 50  0001 C CNN
F 3 "~" H 1600 5850 50  0001 C CNN
	1    1600 5850
	1    0    0    -1  
$EndComp
Connection ~ 1600 5850
Wire Wire Line
	1600 5850 1850 5850
Wire Wire Line
	3550 5850 3500 5850
Wire Wire Line
	3500 6400 3500 5850
Connection ~ 3500 5850
Wire Wire Line
	3800 6600 3800 6900
Wire Wire Line
	3500 6800 3500 6900
Wire Wire Line
	3500 7300 3500 7400
Wire Wire Line
	4200 7300 4200 7400
Wire Wire Line
	4200 7000 4200 6900
Connection ~ 3500 6900
Wire Wire Line
	3500 6900 3500 7000
Wire Wire Line
	3800 6900 3500 6900
Text Label 4350 5850 0    50   ~ 0
P5V
Text Label 3200 1150 0    50   ~ 0
P5V
Wire Wire Line
	1050 5850 1600 5850
$Comp
L dmmt5401:DMMT5401-raspberrypi_hat-powerpiHAT-rescue-powerpiHAT-rescue Q2
U 1 1 5E9DE6E3
P 3600 6600
F 0 "Q2" H 3791 6509 50  0000 L CNN
F 1 "DMMT5401" H 3791 6600 50  0000 L CNN
F 2 "Package_TO_SOT_SMD:SOT-23-6" H 3791 6691 50  0001 L CIN
F 3 "" H 3600 6600 50  0000 L CNN
F 4 "TRANS 2PNP 150V 0.2A SOT26" H 3600 6600 50  0001 C CNN "Description"
F 5 "DMMT5401-FDICT-ND" H 3600 6600 50  0001 C CNN "Digi-Key_PN"
	1    3600 6600
	-1   0    0    1   
$EndComp
$Comp
L dmmt5401:DMMT5401-raspberrypi_hat-powerpiHAT-rescue-powerpiHAT-rescue Q2
U 2 1 5E9DF95B
P 4100 6600
F 0 "Q2" H 4291 6509 50  0000 L CNN
F 1 "DMMT5401" H 4291 6600 50  0000 L CNN
F 2 "Package_TO_SOT_SMD:SOT-23-6" H 4291 6691 50  0001 L CIN
F 3 "" H 4100 6600 50  0000 L CNN
F 4 "TRANS 2PNP 150V 0.2A SOT26" H 4100 6600 50  0001 C CNN "Description"
F 5 "DMMT5401-FDICT-ND" H 4100 6600 50  0001 C CNN "Digi-Key_PN"
	2    4100 6600
	1    0    0    1   
$EndComp
Wire Wire Line
	3900 6600 3900 6900
Wire Wire Line
	3900 6900 3800 6900
Connection ~ 3800 6900
Wire Wire Line
	3950 5850 4200 5850
Wire Wire Line
	4200 6400 4200 5850
Connection ~ 4200 5850
Wire Wire Line
	4200 5850 4350 5850
$Comp
L Device:R R3
U 1 1 5E9FB648
P 7450 5400
F 0 "R3" H 7520 5446 50  0000 L CNN
F 1 "3k" H 7520 5355 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 7380 5400 50  0001 C CNN
F 3 "~" H 7450 5400 50  0001 C CNN
F 4 "RES 3K OHM 0.5% 1/8W 0805" H 7450 5400 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K00CT-ND" H 7450 5400 50  0001 C CNN "Digi-Key_PN"
	1    7450 5400
	1    0    0    -1  
$EndComp
$Comp
L Device:R R5
U 1 1 5E9FD922
P 7800 5400
F 0 "R5" H 7870 5446 50  0000 L CNN
F 1 "3k" H 7870 5355 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 7730 5400 50  0001 C CNN
F 3 "~" H 7800 5400 50  0001 C CNN
F 4 "RES 3K OHM 0.5% 1/8W 0805" H 7800 5400 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K00CT-ND" H 7800 5400 50  0001 C CNN "Digi-Key_PN"
	1    7800 5400
	1    0    0    -1  
$EndComp
$Comp
L Device:R R7
U 1 1 5E9FE04F
P 8150 5400
F 0 "R7" H 8220 5446 50  0000 L CNN
F 1 "3k" H 8220 5355 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 8080 5400 50  0001 C CNN
F 3 "~" H 8150 5400 50  0001 C CNN
F 4 "RES 3K OHM 0.5% 1/8W 0805" H 8150 5400 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K00CT-ND" H 8150 5400 50  0001 C CNN "Digi-Key_PN"
	1    8150 5400
	1    0    0    -1  
$EndComp
$Comp
L Device:R R4
U 1 1 5E9FE6D5
P 7450 5800
F 0 "R4" H 7520 5846 50  0000 L CNN
F 1 "3.9k" H 7520 5755 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 7380 5800 50  0001 C CNN
F 3 "~" H 7450 5800 50  0001 C CNN
F 4 "RES 3.9K OHM 0.5% 1/8W 0805" H 7450 5800 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K90CT-ND" H 7450 5800 50  0001 C CNN "Digi-Key_PN"
	1    7450 5800
	1    0    0    -1  
$EndComp
$Comp
L Device:R R6
U 1 1 5E9FEF4C
P 7800 5800
F 0 "R6" H 7870 5846 50  0000 L CNN
F 1 "3.9k" H 7870 5755 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 7730 5800 50  0001 C CNN
F 3 "~" H 7800 5800 50  0001 C CNN
F 4 "RES 3.9K OHM 0.5% 1/8W 0805" H 7800 5800 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K90CT-ND" H 7800 5800 50  0001 C CNN "Digi-Key_PN"
	1    7800 5800
	1    0    0    -1  
$EndComp
$Comp
L Device:R R8
U 1 1 5E9FF34A
P 8150 5800
F 0 "R8" H 8220 5846 50  0000 L CNN
F 1 "3.9k" H 8220 5755 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 8080 5800 50  0001 C CNN
F 3 "~" H 8150 5800 50  0001 C CNN
F 4 "RES 3.9K OHM 0.5% 1/8W 0805" H 8150 5800 50  0001 C CNN "Description"
F 5 "RNCF0805DTE3K90CT-ND" H 8150 5800 50  0001 C CNN "Digi-Key_PN"
	1    8150 5800
	1    0    0    -1  
$EndComp
Wire Wire Line
	8150 5550 8150 5600
Wire Wire Line
	7800 5550 7800 5600
Wire Wire Line
	7450 5550 7450 5600
Wire Wire Line
	7450 5950 7800 5950
Wire Wire Line
	7800 5950 8150 5950
Connection ~ 7800 5950
$Comp
L power:GND #PWR0110
U 1 1 5EA0E245
P 7800 5950
F 0 "#PWR0110" H 7800 5700 50  0001 C CNN
F 1 "GND" H 7805 5777 50  0000 C CNN
F 2 "" H 7800 5950 50  0001 C CNN
F 3 "" H 7800 5950 50  0001 C CNN
	1    7800 5950
	1    0    0    -1  
$EndComp
Wire Wire Line
	7450 5600 7250 5600
Wire Wire Line
	7250 5600 7250 6050
Wire Wire Line
	7250 6050 7050 6050
Connection ~ 7450 5600
Wire Wire Line
	7450 5600 7450 5650
Wire Wire Line
	7800 5600 7600 5600
Wire Wire Line
	7600 5600 7600 6200
Wire Wire Line
	7600 6200 7050 6200
Connection ~ 7800 5600
Wire Wire Line
	7800 5600 7800 5650
Wire Wire Line
	8150 5600 8000 5600
Wire Wire Line
	8000 5600 8000 6350
Wire Wire Line
	8000 6350 7050 6350
Connection ~ 8150 5600
Wire Wire Line
	8150 5600 8150 5650
Wire Wire Line
	7450 5250 7450 5100
Wire Wire Line
	7800 5250 7800 5100
Wire Wire Line
	8150 5250 8150 5100
Text Label 7450 5100 0    50   ~ 0
ECHO_0
Text Label 7800 5100 0    50   ~ 0
ECHO_1
Text Label 8150 5100 0    50   ~ 0
ECHO_2
Wire Wire Line
	5600 4250 6550 4250
Text Label 5600 4350 0    50   ~ 0
ECHO_0
Text Label 5600 5050 0    50   ~ 0
GPIO5
Text Label 5600 4550 0    50   ~ 0
P5V
Text Label 5600 5150 0    50   ~ 0
P5V
Text Label 5600 5750 0    50   ~ 0
P5V
$Comp
L power:GND #PWR0111
U 1 1 5EA6EE54
P 6550 5800
F 0 "#PWR0111" H 6550 5550 50  0001 C CNN
F 1 "GND" H 6555 5627 50  0000 C CNN
F 2 "" H 6550 5800 50  0001 C CNN
F 3 "" H 6550 5800 50  0001 C CNN
	1    6550 5800
	1    0    0    -1  
$EndComp
Wire Wire Line
	5600 4850 6550 4850
Wire Wire Line
	6550 4850 6550 4250
Wire Wire Line
	6550 4850 6550 5450
Wire Wire Line
	6550 5450 5600 5450
Connection ~ 6550 4850
Wire Wire Line
	6550 5450 6550 5800
Connection ~ 6550 5450
Text Label 5600 4950 0    50   ~ 0
ECHO_1
Text Label 5600 5550 0    50   ~ 0
ECHO_2
Text Label 5600 4450 0    50   ~ 0
GPIO16
Text Label 5600 5650 0    50   ~ 0
GPIO25
Wire Wire Line
	7350 4250 6550 4250
Connection ~ 6550 4250
Text Label 7100 4750 0    50   ~ 0
P5V
Wire Wire Line
	7100 4750 7350 4750
Wire Wire Line
	7350 4650 7100 4650
Wire Wire Line
	7350 4550 7100 4550
Wire Wire Line
	7350 4450 7100 4450
Wire Wire Line
	7350 4350 7100 4350
Text Label 7100 4350 0    50   ~ 0
GPIO19
Text Label 7100 4550 0    50   ~ 0
GPIO20
Text Label 7100 4650 0    50   ~ 0
GPIO21
Text Label 7100 4450 0    50   ~ 0
GPIO26
Text Label 7050 6200 0    50   ~ 0
GPIO6
Wire Wire Line
	2450 5850 2650 5850
Connection ~ 2650 5850
Wire Wire Line
	2650 5850 3500 5850
Wire Wire Line
	2650 6150 2150 6150
Connection ~ 2150 6150
Wire Wire Line
	2150 6150 1600 6150
$Comp
L power:GND #PWR0109
U 1 1 5E98E854
P 2150 6300
F 0 "#PWR0109" H 2150 6050 50  0001 C CNN
F 1 "GND" H 2155 6127 50  0000 C CNN
F 2 "" H 2150 6300 50  0001 C CNN
F 3 "" H 2150 6300 50  0001 C CNN
	1    2150 6300
	1    0    0    -1  
$EndComp
Wire Wire Line
	2150 6150 2150 6300
Wire Wire Line
	4200 6800 4200 6900
Wire Wire Line
	3850 6150 4550 6150
Wire Wire Line
	4550 6150 4550 6900
Wire Wire Line
	4550 6900 4200 6900
Connection ~ 4200 6900
$Comp
L Memory_EEPROM:CAT24C128 U2
U 1 1 5E97B9E0
P 2050 4550
F 0 "U2" H 2050 4069 50  0000 C CNN
F 1 "CAT24C32" H 2050 4160 50  0000 C CNN
F 2 "Package_SO:SOIC-8_3.9x4.9mm_P1.27mm" H 2050 4550 50  0001 C CNN
F 3 "https://www.onsemi.com/pub/Collateral/CAT24C128-D.PDF" H 2050 4550 50  0001 C CNN
F 4 "IC EEPROM 32K I2C 1MHZ 8SOIC" H 2050 4550 50  0001 C CNN "Description"
F 5 "CAT24C32WI-GT3CT-ND" H 2050 4550 50  0001 C CNN "Digi-Key_PN"
	1    2050 4550
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5C778511
P 1450 3250
F 0 "#FLG0102" H 1450 3325 50  0001 C CNN
F 1 "PWR_FLAG" H 1450 3424 50  0000 C CNN
F 2 "" H 1450 3250 50  0001 C CNN
F 3 "~" H 1450 3250 50  0001 C CNN
	1    1450 3250
	1    0    0    -1  
$EndComp
Text Label 1800 1150 0    50   ~ 0
P3V3_HAT
Text Label 2950 3900 0    50   ~ 0
P3V3_HAT
Wire Wire Line
	2050 4850 2050 4950
$Comp
L power:GND #PWR01
U 1 1 5E99B7F6
P 2050 5100
F 0 "#PWR01" H 2050 4850 50  0001 C CNN
F 1 "GND" H 2055 4927 50  0001 C CNN
F 2 "" H 2050 5100 50  0001 C CNN
F 3 "" H 2050 5100 50  0001 C CNN
	1    2050 5100
	1    0    0    -1  
$EndComp
Wire Wire Line
	2050 3900 2050 4250
$Comp
L Device:R R9
U 1 1 5E9B9AF9
P 2700 4050
F 0 "R9" H 2770 4096 50  0000 L CNN
F 1 "10K" H 2770 4005 50  0000 L CNN
F 2 "Resistor_SMD:R_0805_2012Metric" V 2630 4050 50  0001 C CNN
F 3 "~" H 2700 4050 50  0001 C CNN
F 4 "RES 10K OHM 0.5% 1/8W 0805" H 2700 4050 50  0001 C CNN "Description"
F 5 "RNCF0805DTE10K0CT-ND" H 2700 4050 50  0001 C CNN "Digi-Key_PN"
	1    2700 4050
	1    0    0    -1  
$EndComp
Wire Wire Line
	2050 3900 2700 3900
Connection ~ 2700 3900
Wire Wire Line
	2700 3900 2950 3900
Wire Wire Line
	2450 4550 2850 4550
Wire Wire Line
	1650 4450 1650 4550
Wire Wire Line
	1650 4550 1650 4650
Connection ~ 1650 4550
Wire Wire Line
	1650 4650 1650 4950
Wire Wire Line
	1650 4950 2050 4950
Connection ~ 1650 4650
Connection ~ 2050 4950
Wire Wire Line
	2050 4950 2050 5100
Wire Wire Line
	2450 4450 2850 4450
$Comp
L power:GND #PWR0108
U 1 1 5E98491B
P 1050 6150
F 0 "#PWR0108" H 1050 5900 50  0001 C CNN
F 1 "GND" H 1055 5977 50  0000 C CNN
F 2 "" H 1050 6150 50  0001 C CNN
F 3 "" H 1050 6150 50  0001 C CNN
	1    1050 6150
	1    0    0    -1  
$EndComp
Text Label 1000 1650 0    50   ~ 0
GPIO17
Text Label 2850 4650 0    50   ~ 0
GPIO12
Text Label 7050 6050 0    50   ~ 0
GPIO13
Text Label 7050 6350 0    50   ~ 0
GPIO17
Text Label 1050 5850 0    50   ~ 0
P12V
Text Label 2850 4450 0    50   ~ 0
ID_SD
Text Label 2850 4550 0    50   ~ 0
ID_SC
Text Label 1000 2450 0    50   ~ 0
ID_SD
Text Label 3600 2450 0    50   ~ 0
ID_SC
Text Label 1050 5950 0    50   ~ 0
P-12V
Text Notes 1650 3750 0    138  ~ 28
HAT EEPROM
Text Notes 1250 750  0    138  ~ 28
Raspberry Pi Connector\n
Text Notes 1300 5550 0    138  ~ 28
12v Power Source
Text Notes 4900 3950 0    138  ~ 28
Proximity Transponders & Relay Connectors
$Comp
L dk_Terminal-Blocks-Wire-to-Board:1935161 J2
U 1 1 5EB483BA
P 850 5850
F 0 "J2" V 625 5933 50  0000 C CNN
F 1 "1935161" V 716 5933 50  0000 C CNN
F 2 "digikey-footprints:Terminal_Block_D1.3mm_P5mm" H 1050 6050 60  0001 L CNN
F 3 "https://media.digikey.com/pdf/Data%20Sheets/Phoenix%20Contact%20PDFs/1935161.pdf" H 1050 6150 60  0001 L CNN
F 4 "277-1667-ND" H 1050 6250 60  0001 L CNN "Digi-Key_PN"
F 5 "1935161" H 1050 6350 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 1050 6450 60  0001 L CNN "Category"
F 7 "Terminal Blocks - Wire to Board" H 1050 6550 60  0001 L CNN "Family"
F 8 "https://media.digikey.com/pdf/Data%20Sheets/Phoenix%20Contact%20PDFs/1935161.pdf" H 1050 6650 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/phoenix-contact/1935161/277-1667-ND/568614" H 1050 6750 60  0001 L CNN "DK_Detail_Page"
F 10 "TERM BLK 2POS SIDE ENTRY 5MM PCB" H 1050 6850 60  0001 L CNN "Description"
F 11 "Phoenix Contact" H 1050 6950 60  0001 L CNN "Manufacturer"
F 12 "Active" H 1050 7050 60  0001 L CNN "Status"
	1    850  5850
	0    1    1    0   
$EndComp
NoConn ~ 3600 1650
Wire Wire Line
	2450 4650 2700 4650
Wire Wire Line
	2700 4200 2700 4650
Connection ~ 2700 4650
Wire Wire Line
	2700 4650 2850 4650
NoConn ~ 1000 1450
NoConn ~ 3600 1450
NoConn ~ 3600 1550
NoConn ~ 2050 3050
NoConn ~ 2050 2350
Wire Wire Line
	1850 1550 1850 3200
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B4B-PH-K-S_LF__SN_ J3
U 1 1 5EA109B6
P 5500 4250
F 0 "J3" V 5275 4258 50  0000 C CNN
F 1 "B4B-PH-K-S_LF__SN_" V 5366 4258 50  0000 C CNN
F 2 "Connector_JST:JST_PH_B4B-PH-K_1x04_P2.00mm_Vertical" H 5700 4450 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 4550 60  0001 L CNN
F 4 "455-1706-ND" H 5700 4650 60  0001 L CNN "Digi-Key_PN"
F 5 "B4B-PH-K-S(LF)(SN)" H 5700 4750 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 5700 4850 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 5700 4950 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 5050 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 5700 5150 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 4POS 2MM" H 5700 5250 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 5700 5350 60  0001 L CNN "Manufacturer"
F 12 "Active" H 5700 5450 60  0001 L CNN "Status"
	1    5500 4250
	0    1    1    0   
$EndComp
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B4B-PH-K-S_LF__SN_ J4
U 1 1 5EA18FB2
P 5500 4850
F 0 "J4" V 5275 4858 50  0000 C CNN
F 1 "B4B-PH-K-S_LF__SN_" V 5366 4858 50  0000 C CNN
F 2 "Connector_JST:JST_PH_B4B-PH-K_1x04_P2.00mm_Vertical" H 5700 5050 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 5150 60  0001 L CNN
F 4 "455-1706-ND" H 5700 5250 60  0001 L CNN "Digi-Key_PN"
F 5 "B4B-PH-K-S(LF)(SN)" H 5700 5350 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 5700 5450 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 5700 5550 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 5650 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 5700 5750 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 4POS 2MM" H 5700 5850 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 5700 5950 60  0001 L CNN "Manufacturer"
F 12 "Active" H 5700 6050 60  0001 L CNN "Status"
	1    5500 4850
	0    1    1    0   
$EndComp
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B4B-PH-K-S_LF__SN_ J5
U 1 1 5EA19951
P 5500 5450
F 0 "J5" V 5275 5458 50  0000 C CNN
F 1 "B4B-PH-K-S_LF__SN_" V 5366 5458 50  0000 C CNN
F 2 "Connector_JST:JST_PH_B4B-PH-K_1x04_P2.00mm_Vertical" H 5700 5650 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 5750 60  0001 L CNN
F 4 "455-1706-ND" H 5700 5850 60  0001 L CNN "Digi-Key_PN"
F 5 "B4B-PH-K-S(LF)(SN)" H 5700 5950 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 5700 6050 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 5700 6150 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 5700 6250 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 5700 6350 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 4POS 2MM" H 5700 6450 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 5700 6550 60  0001 L CNN "Manufacturer"
F 12 "Active" H 5700 6650 60  0001 L CNN "Status"
	1    5500 5450
	0    1    1    0   
$EndComp
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B6B-PH-K-S_LF__SN_ J6
U 1 1 5EA33896
P 7450 4250
F 0 "J6" V 7649 4122 50  0000 R CNN
F 1 "B6B-PH-K-S_LF__SN_" V 7740 4122 50  0000 R CNN
F 2 "Connector_JST:JST_PH_B6B-PH-K_1x06_P2.00mm_Vertical" H 7650 4450 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7650 4550 60  0001 L CNN
F 4 "455-1708-ND" H 7650 4650 60  0001 L CNN "Digi-Key_PN"
F 5 "B6B-PH-K-S(LF)(SN)" H 7650 4750 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 7650 4850 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 7650 4950 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7650 5050 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 7650 5150 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 6POS 2MM" H 7650 5250 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 7650 5350 60  0001 L CNN "Manufacturer"
F 12 "Active" H 7650 5450 60  0001 L CNN "Status"
	1    7450 4250
	0    -1   1    0   
$EndComp
Wire Wire Line
	7200 2000 7600 2000
Text Notes 5400 1750 0    138  ~ 28
A/D (ADS1115) Connection
Wire Wire Line
	7200 2550 7200 2000
$Comp
L power:GND #PWR02
U 1 1 5EA77922
P 7050 3150
F 0 "#PWR02" H 7050 2900 50  0001 C CNN
F 1 "GND" H 7055 2977 50  0001 C CNN
F 2 "" H 7050 3150 50  0001 C CNN
F 3 "" H 7050 3150 50  0001 C CNN
	1    7050 3150
	1    0    0    -1  
$EndComp
Wire Wire Line
	7050 2350 7050 3150
Connection ~ 7050 2350
Wire Wire Line
	7050 2350 7050 2050
Text Label 6400 2850 0    50   ~ 0
A3
Text Label 6400 2750 0    50   ~ 0
A2
Text Label 6400 2650 0    50   ~ 0
A1
Text Label 6400 2550 0    50   ~ 0
A0
Text Label 6400 2450 0    50   ~ 0
ALERT
Text Label 6400 2050 0    50   ~ 0
GND
Text Label 6400 2350 0    50   ~ 0
ADDR
Text Label 6400 2250 0    50   ~ 0
GPIO2_SDA1
Text Label 6400 2150 0    50   ~ 0
GPIO3_SCL1
Text Label 6400 1950 0    50   ~ 0
P3V3_HAT
Wire Wire Line
	6400 2550 7200 2550
NoConn ~ 6400 2450
Wire Wire Line
	6400 2350 7050 2350
Wire Wire Line
	7050 2050 6400 2050
$Comp
L Connector:Conn_01x10_Male J8
U 1 1 5EA08AF1
P 6200 2350
F 0 "J8" H 6308 2931 50  0000 C CNN
F 1 "TI ADS1115 Breakout" H 6308 2840 50  0000 C CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x10_P2.54mm_Vertical" H 6200 2350 50  0001 C CNN
F 3 "~" H 6200 2350 50  0001 C CNN
	1    6200 2350
	1    0    0    -1  
$EndComp
Text Label 7600 1900 2    50   ~ 0
P3V3_HAT
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B4B-PH-K-S_LF__SN_ J7
U 1 1 5EAAA21F
P 7700 1900
F 0 "J7" V 7799 1772 50  0000 R CNN
F 1 "B4B-PH-K-S_LF__SN_" V 7890 1772 50  0000 R CNN
F 2 "Connector_JST:JST_PH_B4B-PH-K_1x04_P2.00mm_Vertical" H 7900 2100 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7900 2200 60  0001 L CNN
F 4 "455-1706-ND" H 7900 2300 60  0001 L CNN "Digi-Key_PN"
F 5 "B4B-PH-K-S(LF)(SN)" H 7900 2400 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 7900 2500 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 7900 2600 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7900 2700 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 7900 2800 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 4POS 2MM" H 7900 2900 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 7900 3000 60  0001 L CNN "Manufacturer"
F 12 "Active" H 7900 3100 60  0001 L CNN "Status"
	1    7700 1900
	0    -1   1    0   
$EndComp
$Comp
L dk_Rectangular-Connectors-Headers-Male-Pins:B4B-PH-K-S_LF__SN_ J9
U 1 1 5EAABF94
P 7700 2450
F 0 "J9" V 7799 2322 50  0000 R CNN
F 1 "B4B-PH-K-S_LF__SN_" V 7890 2322 50  0000 R CNN
F 2 "Connector_JST:JST_PH_B4B-PH-K_1x04_P2.00mm_Vertical" H 7900 2650 60  0001 L CNN
F 3 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7900 2750 60  0001 L CNN
F 4 "455-1706-ND" H 7900 2850 60  0001 L CNN "Digi-Key_PN"
F 5 "B4B-PH-K-S(LF)(SN)" H 7900 2950 60  0001 L CNN "MPN"
F 6 "Connectors, Interconnects" H 7900 3050 60  0001 L CNN "Category"
F 7 "Rectangular Connectors - Headers, Male Pins" H 7900 3150 60  0001 L CNN "Family"
F 8 "http://www.jst-mfg.com/product/pdf/eng/ePH.pdf" H 7900 3250 60  0001 L CNN "DK_Datasheet_Link"
F 9 "/product-detail/en/jst-sales-america-inc/B4B-PH-K-S(LF)(SN)/455-1706-ND/926613" H 7900 3350 60  0001 L CNN "DK_Detail_Page"
F 10 "CONN HEADER VERT 4POS 2MM" H 7900 3450 60  0001 L CNN "Description"
F 11 "JST Sales America Inc." H 7900 3550 60  0001 L CNN "Manufacturer"
F 12 "Active" H 7900 3650 60  0001 L CNN "Status"
	1    7700 2450
	0    -1   1    0   
$EndComp
Text Label 7600 2450 2    50   ~ 0
P3V3_HAT
Text Label 7600 2750 2    50   ~ 0
GND
Text Label 7600 2200 2    50   ~ 0
GND
Wire Wire Line
	7600 2100 7300 2100
Wire Wire Line
	7300 2100 7300 2650
Wire Wire Line
	7300 2650 6400 2650
Wire Wire Line
	6400 2750 7400 2750
Wire Wire Line
	7400 2750 7400 2550
Wire Wire Line
	7400 2550 7600 2550
Wire Wire Line
	7600 2650 7500 2650
Wire Wire Line
	7500 2650 7500 2850
Wire Wire Line
	7500 2850 6400 2850
$EndSCHEMATC