package irsdk

type Session struct {
	WeekendInfo struct {
		TrackName              string `yaml:"TrackName"`
		TrackID                int    `yaml:"TrackID"`
		TrackLength            string `yaml:"TrackLength"`
		TrackLengthOfficial    string `yaml:"TrackLengthOfficial"`
		TrackDisplayName       string `yaml:"TrackDisplayName"`
		TrackDisplayShortName  string `yaml:"TrackDisplayShortName"`
		TrackConfigName        string `yaml:"TrackConfigName"`
		TrackCity              string `yaml:"TrackCity"`
		TrackCountry           string `yaml:"TrackCountry"`
		TrackAltitude          string `yaml:"TrackAltitude"`
		TrackLatitude          string `yaml:"TrackLatitude"`
		TrackLongitude         string `yaml:"TrackLongitude"`
		TrackNorthOffset       string `yaml:"TrackNorthOffset"`
		TrackNumTurns          int    `yaml:"TrackNumTurns"`
		TrackPitSpeedLimit     string `yaml:"TrackPitSpeedLimit"`
		TrackType              string `yaml:"TrackType"`
		TrackDirection         string `yaml:"TrackDirection"`
		TrackWeatherType       string `yaml:"TrackWeatherType"`
		TrackSkies             string `yaml:"TrackSkies"`
		TrackSurfaceTemp       string `yaml:"TrackSurfaceTemp"`
		TrackAirTemp           string `yaml:"TrackAirTemp"`
		TrackAirPressure       string `yaml:"TrackAirPressure"`
		TrackWindVel           string `yaml:"TrackWindVel"`
		TrackWindDir           string `yaml:"TrackWindDir"`
		TrackRelativeHumidity  string `yaml:"TrackRelativeHumidity"`
		TrackFogLevel          string `yaml:"TrackFogLevel"`
		TrackPrecipitation     string `yaml:"TrackPrecipitation"`
		TrackCleanup           int    `yaml:"TrackCleanup"`
		TrackDynamicTrack      int    `yaml:"TrackDynamicTrack"`
		TrackVersion           string `yaml:"TrackVersion"`
		SeriesID               int    `yaml:"SeriesID"`
		SeasonID               int    `yaml:"SeasonID"`
		SessionID              int    `yaml:"SessionID"`
		SubSessionID           int    `yaml:"SubSessionID"`
		LeagueID               int    `yaml:"LeagueID"`
		Official               int    `yaml:"Official"`
		RaceWeek               int    `yaml:"RaceWeek"`
		EventType              string `yaml:"EventType"`
		Category               string `yaml:"Category"`
		SimMode                string `yaml:"SimMode"`
		TeamRacing             int    `yaml:"TeamRacing"`
		MinDrivers             int    `yaml:"MinDrivers"`
		MaxDrivers             int    `yaml:"MaxDrivers"`
		DCRuleSet              string `yaml:"DCRuleSet"`
		QualifierMustStartRace int    `yaml:"QualifierMustStartRace"`
		NumCarClasses          int    `yaml:"NumCarClasses"`
		NumCarTypes            int    `yaml:"NumCarTypes"`
		HeatRacing             int    `yaml:"HeatRacing"`
		BuildType              string `yaml:"BuildType"`
		BuildTarget            string `yaml:"BuildTarget"`
		BuildVersion           string `yaml:"BuildVersion"`
		WeekendOptions         struct {
			NumStarters      int    `yaml:"NumStarters"`
			StartingGrid     string `yaml:"StartingGrid"`
			QualifyScoring   string `yaml:"QualifyScoring"`
			CourseCautions   string `yaml:"CourseCautions"`
			StandingStart    int    `yaml:"StandingStart"`
			ShortParadeLap   int    `yaml:"ShortParadeLap"`
			Restarts         string `yaml:"Restarts"`
			WeatherType      string `yaml:"WeatherType"`
			Skies            string `yaml:"Skies"`
			WindDirection    string `yaml:"WindDirection"`
			WindSpeed        string `yaml:"WindSpeed"`
			WeatherTemp      string `yaml:"WeatherTemp"`
			RelativeHumidity string `yaml:"RelativeHumidity"`
			FogLevel         string `yaml:"FogLevel"`
			TimeOfDay        string `yaml:"TimeOfDay"`
			Date             struct {
			} `yaml:"Date"`
			EarthRotationSpeedupFactor int    `yaml:"EarthRotationSpeedupFactor"`
			Unofficial                 int    `yaml:"Unofficial"`
			CommercialMode             string `yaml:"CommercialMode"`
			NightMode                  string `yaml:"NightMode"`
			IsFixedSetup               int    `yaml:"IsFixedSetup"`
			StrictLapsChecking         string `yaml:"StrictLapsChecking"`
			HasOpenRegistration        int    `yaml:"HasOpenRegistration"`
			HardcoreLevel              int    `yaml:"HardcoreLevel"`
			NumJokerLaps               int    `yaml:"NumJokerLaps"`
			IncidentLimit              string `yaml:"IncidentLimit"`
			FastRepairsLimit           int    `yaml:"FastRepairsLimit"`
			GreenWhiteCheckeredLimit   int    `yaml:"GreenWhiteCheckeredLimit"`
		} `yaml:"WeekendOptions"`
		TelemetryOptions struct {
			TelemetryDiskFile string `yaml:"TelemetryDiskFile"`
		} `yaml:"TelemetryOptions"`
	} `yaml:"WeekendInfo"`
	SessionInfo struct {
		Sessions []struct {
			SessionNum                       int         `yaml:"SessionNum"`
			SessionLaps                      string      `yaml:"SessionLaps"`
			SessionTime                      string      `yaml:"SessionTime"`
			SessionNumLapsToAvg              int         `yaml:"SessionNumLapsToAvg"`
			SessionType                      string      `yaml:"SessionType"`
			SessionTrackRubberState          string      `yaml:"SessionTrackRubberState"`
			SessionName                      string      `yaml:"SessionName"`
			SessionSubType                   interface{} `yaml:"SessionSubType"`
			SessionSkipped                   int         `yaml:"SessionSkipped"`
			SessionRunGroupsUsed             int         `yaml:"SessionRunGroupsUsed"`
			SessionEnforceTireCompoundChange int         `yaml:"SessionEnforceTireCompoundChange"`
			ResultsPositions                 []struct {
				Position          int     `yaml:"Position"`
				ClassPosition     int     `yaml:"ClassPosition"`
				CarIdx            int     `yaml:"CarIdx"`
				Lap               int     `yaml:"Lap"`
				Time              float64 `yaml:"Time"`
				FastestLap        int     `yaml:"FastestLap"`
				FastestTime       float64 `yaml:"FastestTime"`
				LastTime          float64 `yaml:"LastTime"`
				LapsLed           int     `yaml:"LapsLed"`
				LapsComplete      int     `yaml:"LapsComplete"`
				JokerLapsComplete int     `yaml:"JokerLapsComplete"`
				LapsDriven        float64 `yaml:"LapsDriven"`
				Incidents         int     `yaml:"Incidents"`
				ReasonOutID       int     `yaml:"ReasonOutId"`
				ReasonOutStr      string  `yaml:"ReasonOutStr"`
			} `yaml:"ResultsPositions"`
			ResultsFastestLap []struct {
				CarIdx      int     `yaml:"CarIdx"`
				FastestLap  int     `yaml:"FastestLap"`
				FastestTime float64 `yaml:"FastestTime"`
			} `yaml:"ResultsFastestLap"`
			ResultsAverageLapTime  int `yaml:"ResultsAverageLapTime"`
			ResultsNumCautionFlags int `yaml:"ResultsNumCautionFlags"`
			ResultsNumCautionLaps  int `yaml:"ResultsNumCautionLaps"`
			ResultsNumLeadChanges  int `yaml:"ResultsNumLeadChanges"`
			ResultsLapsComplete    int `yaml:"ResultsLapsComplete"`
			ResultsOfficial        int `yaml:"ResultsOfficial"`
		} `yaml:"Sessions"`
	} `yaml:"SessionInfo"`
	CameraInfo struct {
		Groups []struct {
			GroupNum  int    `yaml:"GroupNum"`
			GroupName string `yaml:"GroupName"`
			Cameras   []struct {
				CameraNum  int    `yaml:"CameraNum"`
				CameraName string `yaml:"CameraName"`
			} `yaml:"Cameras"`
			IsScenic bool `yaml:"IsScenic,omitempty"`
		} `yaml:"Groups"`
	} `yaml:"CameraInfo"`
	RadioInfo struct {
		SelectedRadioNum int `yaml:"SelectedRadioNum"`
		Radios           []struct {
			RadioNum            int `yaml:"RadioNum"`
			HopCount            int `yaml:"HopCount"`
			NumFrequencies      int `yaml:"NumFrequencies"`
			TunedToFrequencyNum int `yaml:"TunedToFrequencyNum"`
			ScanningIsOn        int `yaml:"ScanningIsOn"`
			Frequencies         []struct {
				FrequencyNum  int    `yaml:"FrequencyNum"`
				FrequencyName string `yaml:"FrequencyName"`
				Priority      int    `yaml:"Priority"`
				CarIdx        int    `yaml:"CarIdx"`
				EntryIdx      int    `yaml:"EntryIdx"`
				ClubID        int    `yaml:"ClubID"`
				CanScan       int    `yaml:"CanScan"`
				CanSquawk     int    `yaml:"CanSquawk"`
				Muted         int    `yaml:"Muted"`
				IsMutable     int    `yaml:"IsMutable"`
				IsDeletable   int    `yaml:"IsDeletable"`
			} `yaml:"Frequencies"`
		} `yaml:"Radios"`
	} `yaml:"RadioInfo"`
	DriverInfo struct {
		DriverCarIdx              int     `yaml:"DriverCarIdx"`
		DriverUserID              int     `yaml:"DriverUserID"`
		PaceCarIdx                int     `yaml:"PaceCarIdx"`
		DriverHeadPosX            float64 `yaml:"DriverHeadPosX"`
		DriverHeadPosY            float64 `yaml:"DriverHeadPosY"`
		DriverHeadPosZ            float64 `yaml:"DriverHeadPosZ"`
		DriverCarIsElectric       int     `yaml:"DriverCarIsElectric"`
		DriverCarIdleRPM          float64 `yaml:"DriverCarIdleRPM"`
		DriverCarRedLine          float64 `yaml:"DriverCarRedLine"`
		DriverCarEngCylinderCount int     `yaml:"DriverCarEngCylinderCount"`
		DriverCarFuelKgPerLtr     float64 `yaml:"DriverCarFuelKgPerLtr"`
		DriverCarFuelMaxLtr       float64 `yaml:"DriverCarFuelMaxLtr"`
		DriverCarMaxFuelPct       float64 `yaml:"DriverCarMaxFuelPct"`
		DriverCarGearNumForward   int     `yaml:"DriverCarGearNumForward"`
		DriverCarGearNeutral      int     `yaml:"DriverCarGearNeutral"`
		DriverCarGearReverse      int     `yaml:"DriverCarGearReverse"`
		DriverCarSLFirstRPM       float64 `yaml:"DriverCarSLFirstRPM"`
		DriverCarSLShiftRPM       float64 `yaml:"DriverCarSLShiftRPM"`
		DriverCarSLLastRPM        float64 `yaml:"DriverCarSLLastRPM"`
		DriverCarSLBlinkRPM       float64 `yaml:"DriverCarSLBlinkRPM"`
		DriverCarVersion          string  `yaml:"DriverCarVersion"`
		DriverPitTrkPct           float64 `yaml:"DriverPitTrkPct"`
		DriverCarEstLapTime       float64 `yaml:"DriverCarEstLapTime"`
		DriverSetupName           string  `yaml:"DriverSetupName"`
		DriverSetupIsModified     int     `yaml:"DriverSetupIsModified"`
		DriverSetupLoadTypeName   string  `yaml:"DriverSetupLoadTypeName"`
		DriverSetupPassedTech     int     `yaml:"DriverSetupPassedTech"`
		DriverIncidentCount       int     `yaml:"DriverIncidentCount"`
		Drivers                   []struct {
			CarIdx                  int     `yaml:"CarIdx"`
			UserName                string  `yaml:"UserName"`
			AbbrevName              string  `yaml:"AbbrevName"`
			Initials                string  `yaml:"Initials"`
			UserID                  int     `yaml:"UserID"`
			TeamID                  int     `yaml:"TeamID"`
			TeamName                string  `yaml:"TeamName"`
			CarNumber               string  `yaml:"CarNumber"`
			CarNumberRaw            int     `yaml:"CarNumberRaw"`
			CarPath                 string  `yaml:"CarPath"`
			CarClassID              int     `yaml:"CarClassID"`
			CarID                   int     `yaml:"CarID"`
			CarIsPaceCar            int     `yaml:"CarIsPaceCar"`
			CarIsAI                 int     `yaml:"CarIsAI"`
			CarIsElectric           int     `yaml:"CarIsElectric"`
			CarScreenName           string  `yaml:"CarScreenName"`
			CarScreenNameShort      string  `yaml:"CarScreenNameShort"`
			CarClassShortName       string  `yaml:"CarClassShortName"`
			CarClassRelSpeed        int     `yaml:"CarClassRelSpeed"`
			CarClassLicenseLevel    int     `yaml:"CarClassLicenseLevel"`
			CarClassMaxFuelPct      string  `yaml:"CarClassMaxFuelPct"`
			CarClassWeightPenalty   string  `yaml:"CarClassWeightPenalty"`
			CarClassPowerAdjust     string  `yaml:"CarClassPowerAdjust"`
			CarClassDryTireSetLimit string  `yaml:"CarClassDryTireSetLimit"`
			CarClassColor           int     `yaml:"CarClassColor"`
			CarClassEstLapTime      float64 `yaml:"CarClassEstLapTime"`
			IRating                 int     `yaml:"IRating"`
			LicLevel                int     `yaml:"LicLevel"`
			LicSubLevel             int     `yaml:"LicSubLevel"`
			LicString               string  `yaml:"LicString"`
			LicColor                int     `yaml:"LicColor"`
			IsSpectator             int     `yaml:"IsSpectator"`
			CarDesignStr            string  `yaml:"CarDesignStr"`
			HelmetDesignStr         string  `yaml:"HelmetDesignStr"`
			SuitDesignStr           string  `yaml:"SuitDesignStr"`
			BodyType                int     `yaml:"BodyType"`
			FaceType                int     `yaml:"FaceType"`
			HelmetType              int     `yaml:"HelmetType"`
			CarNumberDesignStr      string  `yaml:"CarNumberDesignStr"`
			CarSponsor1             int     `yaml:"CarSponsor_1"`
			CarSponsor2             int     `yaml:"CarSponsor_2"`
			ClubName                string  `yaml:"ClubName"`
			ClubID                  int     `yaml:"ClubID"`
			DivisionName            string  `yaml:"DivisionName"`
			DivisionID              int     `yaml:"DivisionID"`
			CurDriverIncidentCount  int     `yaml:"CurDriverIncidentCount"`
			TeamIncidentCount       int     `yaml:"TeamIncidentCount"`
		} `yaml:"Drivers"`
	} `yaml:"DriverInfo"`
	SplitTimeInfo struct {
		Sectors []struct {
			SectorNum      int     `yaml:"SectorNum"`
			SectorStartPct float64 `yaml:"SectorStartPct"`
		} `yaml:"Sectors"`
	} `yaml:"SplitTimeInfo"`
	CarSetup struct {
		UpdateCount int `yaml:"UpdateCount"`
		TiresAero   struct {
			LeftFront struct {
				StartingPressure string `yaml:"StartingPressure"`
				LastHotPressure  string `yaml:"LastHotPressure"`
				LastTempsOMI     string `yaml:"LastTempsOMI"`
				TreadRemaining   string `yaml:"TreadRemaining"`
			} `yaml:"LeftFront"`
			LeftRear struct {
				StartingPressure string `yaml:"StartingPressure"`
				LastHotPressure  string `yaml:"LastHotPressure"`
				LastTempsOMI     string `yaml:"LastTempsOMI"`
				TreadRemaining   string `yaml:"TreadRemaining"`
			} `yaml:"LeftRear"`
			RightFront struct {
				StartingPressure string `yaml:"StartingPressure"`
				LastHotPressure  string `yaml:"LastHotPressure"`
				LastTempsIMO     string `yaml:"LastTempsIMO"`
				TreadRemaining   string `yaml:"TreadRemaining"`
			} `yaml:"RightFront"`
			RightRear struct {
				StartingPressure string `yaml:"StartingPressure"`
				LastHotPressure  string `yaml:"LastHotPressure"`
				LastTempsIMO     string `yaml:"LastTempsIMO"`
				TreadRemaining   string `yaml:"TreadRemaining"`
			} `yaml:"RightRear"`
			AeroBalanceCalc struct {
				FrontRhAtSpeed string `yaml:"FrontRhAtSpeed"`
				RearRhAtSpeed  string `yaml:"RearRhAtSpeed"`
				WingSetting    string `yaml:"WingSetting"`
				FrontDownforce string `yaml:"FrontDownforce"`
			} `yaml:"AeroBalanceCalc"`
		} `yaml:"TiresAero"`
		Chassis struct {
			Front struct {
				FarbSetting    int    `yaml:"FarbSetting"`
				TotalToeIn     string `yaml:"TotalToeIn"`
				FuelLevel      string `yaml:"FuelLevel"`
				CrossWeight    string `yaml:"CrossWeight"`
				FrontMasterCyl string `yaml:"FrontMasterCyl"`
				RearMasterCyl  string `yaml:"RearMasterCyl"`
				BrakePads      string `yaml:"BrakePads"`
			} `yaml:"Front"`
			LeftFront struct {
				CornerWeight      string `yaml:"CornerWeight"`
				RideHeight        string `yaml:"RideHeight"`
				SpringPerchOffset string `yaml:"SpringPerchOffset"`
				SpringRate        string `yaml:"SpringRate"`
				Camber            string `yaml:"Camber"`
			} `yaml:"LeftFront"`
			LeftRear struct {
				CornerWeight      string `yaml:"CornerWeight"`
				RideHeight        string `yaml:"RideHeight"`
				SpringPerchOffset string `yaml:"SpringPerchOffset"`
				SpringRate        string `yaml:"SpringRate"`
				Camber            string `yaml:"Camber"`
			} `yaml:"LeftRear"`
			Rear struct {
				RarbSetting int    `yaml:"RarbSetting"`
				TotalToeIn  string `yaml:"TotalToeIn"`
				WingSetting string `yaml:"WingSetting"`
			} `yaml:"Rear"`
			InCarDials struct {
				DisplayPage            string `yaml:"DisplayPage"`
				BrakePressureBias      string `yaml:"BrakePressureBias"`
				TractionControlSetting string `yaml:"TractionControlSetting"`
				AbsSetting             string `yaml:"AbsSetting"`
				ThrottleMapSetting     int    `yaml:"ThrottleMapSetting"`
				NightLedStrips         string `yaml:"NightLedStrips"`
			} `yaml:"InCarDials"`
			RightFront struct {
				CornerWeight      string `yaml:"CornerWeight"`
				RideHeight        string `yaml:"RideHeight"`
				SpringPerchOffset string `yaml:"SpringPerchOffset"`
				SpringRate        string `yaml:"SpringRate"`
				Camber            string `yaml:"Camber"`
			} `yaml:"RightFront"`
			RightRear struct {
				CornerWeight      string `yaml:"CornerWeight"`
				RideHeight        string `yaml:"RideHeight"`
				SpringPerchOffset string `yaml:"SpringPerchOffset"`
				SpringRate        string `yaml:"SpringRate"`
				Camber            string `yaml:"Camber"`
			} `yaml:"RightRear"`
			GearsDifferential struct {
				GearStack     string `yaml:"GearStack"`
				FrictionFaces int    `yaml:"FrictionFaces"`
				DiffPreload   string `yaml:"DiffPreload"`
			} `yaml:"GearsDifferential"`
		} `yaml:"Chassis"`
		Dampers struct {
			FrontDampers struct {
				LowSpeedCompressionDamping  string `yaml:"LowSpeedCompressionDamping"`
				HighSpeedCompressionDamping string `yaml:"HighSpeedCompressionDamping"`
				LowSpeedReboundDamping      string `yaml:"LowSpeedReboundDamping"`
				HighSpeedReboundDamping     string `yaml:"HighSpeedReboundDamping"`
			} `yaml:"FrontDampers"`
			RearDampers struct {
				LowSpeedCompressionDamping  string `yaml:"LowSpeedCompressionDamping"`
				HighSpeedCompressionDamping string `yaml:"HighSpeedCompressionDamping"`
				LowSpeedReboundDamping      string `yaml:"LowSpeedReboundDamping"`
				HighSpeedReboundDamping     string `yaml:"HighSpeedReboundDamping"`
			} `yaml:"RearDampers"`
		} `yaml:"Dampers"`
	} `yaml:"CarSetup"`
}
