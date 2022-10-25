<div align="center">
  
# `🌦 TempLI ( Temperature + CLI )`

## **An app made to get weather information through the terminal**

<p align="center">
  <img width=40% height=40% src="https://github.com/Amad3eu/TempLI/blob/main/.github/assets/cloud.gif?raw=true" alt="cloud">
</p>

[![Status Badge](https://github.com/Amad3eu/TempLI/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/Amad3eu/TempLI/actions/workflows/build.yml)

</div>

## Features 🏆

* Can show
  * Population
  * Latitude
  * Longitude
  * Weather
    * Temperature
    * Humidity
    * Wind Speed
    * Wind Direction
    * Weather Conditions
    * Feels Like Temperature (Apparent temperature based on a combination of factors e.g. humidity, wind speed, solar radiation, wind chill and air pressure)
    * Surface Pressure
    * Sealevel Pressure
    * UV Index
* Raw JSON output

## Building 🔨

* If you don't want to manually build the app, you can use:
  <!-- * The [weather-cli]() AUR if you're on Arch -->
  * Grab a binary from the [Artifacts](https://github.com/Amad3eu/TempLI/actions/runs/2861757821)

Otherwise, you can build the program yourself by running:

***Make sure you have Go installed***

```bash
go build
```

* **The executable will be located in the same directory you ran the command from**

## Usage 💻

* Example 1:

```bash
./weather-Cli get Joinville
__________________________
Searching for city Joinville...

City/Country: Joinville/Brazil
Latitude:  -26.3051
Longitude: -48.8461
Timezone: Santa Catarina/Brazil
Population: 597,658 

Weather Info:
        Temperature: 28.5°
        Wind Direction: 168°
        Wind Speed: 7.0 Km/h
        Weather Condition: Partly Cloudy
        Humidity: 50.00%
        Real Feel: 39.1°
        Surface Pressure: 999.60 hPa
        Sealevel Pressure: 1004.00 hPa
        UV Index: 6
```

* Example 2:

```bash
./weather-Cli get Tokyo --raw
__________________________
(Returns the same data but in JSON)
```

## Acknowledgments and FAQ 📚

API used: [Open-Meteo](https://open-meteo.com/)

Built with: [Cobra](https://github.com/spf13/cobra)
