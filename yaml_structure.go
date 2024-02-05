package irsdk

type Session struct {
	WeekendInfo struct {
		TrackName              string `yaml:"TrackName" json:"track_name,omitempty"`
		TrackID                int    `yaml:"TrackID" json:"track_id,omitempty"`
		TrackLength            string `yaml:"TrackLength" json:"track_length,omitempty"`
		TrackLengthOfficial    string `yaml:"TrackLengthOfficial" json:"track_length_official,omitempty"`
		TrackDisplayName       string `yaml:"TrackDisplayName" json:"track_display_name,omitempty"`
		TrackDisplayShortName  string `yaml:"TrackDisplayShortName" json:"track_display_short_name,omitempty"`
		TrackConfigName        string `yaml:"TrackConfigName" json:"track_config_name,omitempty"`
		TrackCity              string `yaml:"TrackCity" json:"track_city,omitempty"`
		TrackCountry           string `yaml:"TrackCountry" json:"track_country,omitempty"`
		TrackAltitude          string `yaml:"TrackAltitude" json:"track_altitude,omitempty"`
		TrackLatitude          string `yaml:"TrackLatitude" json:"track_latitude,omitempty"`
		TrackLongitude         string `yaml:"TrackLongitude" json:"track_longitude,omitempty"`
		TrackNorthOffset       string `yaml:"TrackNorthOffset" json:"track_north_offset,omitempty"`
		TrackNumTurns          int    `yaml:"TrackNumTurns" json:"track_num_turns,omitempty"`
		TrackPitSpeedLimit     string `yaml:"TrackPitSpeedLimit" json:"track_pit_speed_limit,omitempty"`
		TrackType              string `yaml:"TrackType" json:"track_type,omitempty"`
		TrackDirection         string `yaml:"TrackDirection" json:"track_direction,omitempty"`
		TrackWeatherType       string `yaml:"TrackWeatherType" json:"track_weather_type,omitempty"`
		TrackSkies             string `yaml:"TrackSkies" json:"track_skies,omitempty"`
		TrackSurfaceTemp       string `yaml:"TrackSurfaceTemp" json:"track_surface_temp,omitempty"`
		TrackAirTemp           string `yaml:"TrackAirTemp" json:"track_air_temp,omitempty"`
		TrackAirPressure       string `yaml:"TrackAirPressure" json:"track_air_pressure,omitempty"`
		TrackWindVel           string `yaml:"TrackWindVel" json:"track_wind_vel,omitempty"`
		TrackWindDir           string `yaml:"TrackWindDir" json:"track_wind_dir,omitempty"`
		TrackRelativeHumidity  string `yaml:"TrackRelativeHumidity" json:"track_relative_humidity,omitempty"`
		TrackFogLevel          string `yaml:"TrackFogLevel" json:"track_fog_level,omitempty"`
		TrackPrecipitation     string `yaml:"TrackPrecipitation" json:"track_precipitation,omitempty"`
		TrackCleanup           int    `yaml:"TrackCleanup" json:"track_cleanup,omitempty"`
		TrackDynamicTrack      int    `yaml:"TrackDynamicTrack" json:"track_dynamic_track,omitempty"`
		TrackVersion           string `yaml:"TrackVersion" json:"track_version,omitempty"`
		SeriesID               int    `yaml:"SeriesID" json:"series_id,omitempty"`
		SeasonID               int    `yaml:"SeasonID" json:"season_id,omitempty"`
		SessionID              int    `yaml:"SessionID" json:"session_id,omitempty"`
		SubSessionID           int    `yaml:"SubSessionID" json:"sub_session_id,omitempty"`
		LeagueID               int    `yaml:"LeagueID" json:"league_id,omitempty"`
		Official               int    `yaml:"Official" json:"official,omitempty"`
		RaceWeek               int    `yaml:"RaceWeek" json:"race_week,omitempty"`
		EventType              string `yaml:"EventType" json:"event_type,omitempty"`
		Category               string `yaml:"Category" json:"category,omitempty"`
		SimMode                string `yaml:"SimMode" json:"sim_mode,omitempty"`
		TeamRacing             int    `yaml:"TeamRacing" json:"team_racing,omitempty"`
		MinDrivers             int    `yaml:"MinDrivers" json:"min_drivers,omitempty"`
		MaxDrivers             int    `yaml:"MaxDrivers" json:"max_drivers,omitempty"`
		DCRuleSet              string `yaml:"DCRuleSet" json:"dc_rule_set,omitempty"`
		QualifierMustStartRace int    `yaml:"QualifierMustStartRace" json:"qualifier_must_start_race,omitempty"`
		NumCarClasses          int    `yaml:"NumCarClasses" json:"num_car_classes,omitempty"`
		NumCarTypes            int    `yaml:"NumCarTypes" json:"num_car_types,omitempty"`
		HeatRacing             int    `yaml:"HeatRacing" json:"heat_racing,omitempty"`
		BuildType              string `yaml:"BuildType" json:"build_type,omitempty"`
		BuildTarget            string `yaml:"BuildTarget" json:"build_target,omitempty"`
		BuildVersion           string `yaml:"BuildVersion" json:"build_version,omitempty"`
		WeekendOptions         struct {
			NumStarters      int    `yaml:"NumStarters" json:"num_starters,omitempty"`
			StartingGrid     string `yaml:"StartingGrid" json:"starting_grid,omitempty"`
			QualifyScoring   string `yaml:"QualifyScoring" json:"qualify_scoring,omitempty"`
			CourseCautions   string `yaml:"CourseCautions" json:"course_cautions,omitempty"`
			StandingStart    int    `yaml:"StandingStart" json:"standing_start,omitempty"`
			ShortParadeLap   int    `yaml:"ShortParadeLap" json:"short_parade_lap,omitempty"`
			Restarts         string `yaml:"Restarts" json:"restarts,omitempty"`
			WeatherType      string `yaml:"WeatherType" json:"weather_type,omitempty"`
			Skies            string `yaml:"Skies" json:"skies,omitempty"`
			WindDirection    string `yaml:"WindDirection" json:"wind_direction,omitempty"`
			WindSpeed        string `yaml:"WindSpeed" json:"wind_speed,omitempty"`
			WeatherTemp      string `yaml:"WeatherTemp" json:"weather_temp,omitempty"`
			RelativeHumidity string `yaml:"RelativeHumidity" json:"relative_humidity,omitempty"`
			FogLevel         string `yaml:"FogLevel" json:"fog_level,omitempty"`
			TimeOfDay        string `yaml:"TimeOfDay" json:"time_of_day,omitempty"`
			Date             struct {
			} `yaml:"Date" json:"date"`
			EarthRotationSpeedupFactor int    `yaml:"EarthRotationSpeedupFactor" json:"earth_rotation_speedup_factor,omitempty"`
			Unofficial                 int    `yaml:"Unofficial" json:"unofficial,omitempty"`
			CommercialMode             string `yaml:"CommercialMode" json:"commercial_mode,omitempty"`
			NightMode                  string `yaml:"NightMode" json:"night_mode,omitempty"`
			IsFixedSetup               int    `yaml:"IsFixedSetup" json:"is_fixed_setup,omitempty"`
			StrictLapsChecking         string `yaml:"StrictLapsChecking" json:"strict_laps_checking,omitempty"`
			HasOpenRegistration        int    `yaml:"HasOpenRegistration" json:"has_open_registration,omitempty"`
			HardcoreLevel              int    `yaml:"HardcoreLevel" json:"hardcore_level,omitempty"`
			NumJokerLaps               int    `yaml:"NumJokerLaps" json:"num_joker_laps,omitempty"`
			IncidentLimit              string `yaml:"IncidentLimit" json:"incident_limit,omitempty"`
			FastRepairsLimit           int    `yaml:"FastRepairsLimit" json:"fast_repairs_limit,omitempty"`
			GreenWhiteCheckeredLimit   int    `yaml:"GreenWhiteCheckeredLimit" json:"green_white_checkered_limit,omitempty"`
		} `yaml:"WeekendOptions" json:"weekend_options"`
		TelemetryOptions struct {
			TelemetryDiskFile string `yaml:"TelemetryDiskFile" json:"telemetry_disk_file,omitempty"`
		} `yaml:"TelemetryOptions" json:"telemetry_options"`
	} `yaml:"WeekendInfo" json:"weekend_info"`
	SessionInfo struct {
		Sessions []struct {
			SessionNum                       int                `yaml:"SessionNum" json:"session_num,omitempty"`
			SessionLaps                      string             `yaml:"SessionLaps" json:"session_laps,omitempty"`
			SessionTime                      string             `yaml:"SessionTime" json:"session_time,omitempty"`
			SessionNumLapsToAvg              int                `yaml:"SessionNumLapsToAvg" json:"session_num_laps_to_avg,omitempty"`
			SessionType                      string             `yaml:"SessionType" json:"session_type,omitempty"`
			SessionTrackRubberState          string             `yaml:"SessionTrackRubberState" json:"session_track_rubber_state,omitempty"`
			SessionName                      string             `yaml:"SessionName" json:"session_name,omitempty"`
			SessionSubType                   interface{}        `yaml:"SessionSubType" json:"session_sub_type,omitempty"`
			SessionSkipped                   int                `yaml:"SessionSkipped" json:"session_skipped,omitempty"`
			SessionRunGroupsUsed             int                `yaml:"SessionRunGroupsUsed" json:"session_run_groups_used,omitempty"`
			SessionEnforceTireCompoundChange int                `yaml:"SessionEnforceTireCompoundChange" json:"session_enforce_tire_compound_change,omitempty"`
			ResultsPositions                 []ResultsPositions `yaml:"ResultsPositions" json:"results_positions,omitempty"`
			ResultsFastestLap                []struct {
				CarIdx      int     `yaml:"CarIdx" json:"car_idx,omitempty"`
				FastestLap  int     `yaml:"FastestLap" json:"fastest_lap,omitempty"`
				FastestTime float64 `yaml:"FastestTime" json:"fastest_time,omitempty"`
			} `yaml:"ResultsFastestLap" json:"results_fastest_lap,omitempty"`
			ResultsAverageLapTime  int `yaml:"ResultsAverageLapTime" json:"results_average_lap_time,omitempty"`
			ResultsNumCautionFlags int `yaml:"ResultsNumCautionFlags" json:"results_num_caution_flags,omitempty"`
			ResultsNumCautionLaps  int `yaml:"ResultsNumCautionLaps" json:"results_num_caution_laps,omitempty"`
			ResultsNumLeadChanges  int `yaml:"ResultsNumLeadChanges" json:"results_num_lead_changes,omitempty"`
			ResultsLapsComplete    int `yaml:"ResultsLapsComplete" json:"results_laps_complete,omitempty"`
			ResultsOfficial        int `yaml:"ResultsOfficial" json:"results_official,omitempty"`
		} `yaml:"Sessions" json:"sessions,omitempty"`
	} `yaml:"SessionInfo" json:"session_info"`
	CameraInfo struct {
		Groups []struct {
			GroupNum  int    `yaml:"GroupNum" json:"group_num,omitempty"`
			GroupName string `yaml:"GroupName" json:"group_name,omitempty"`
			Cameras   []struct {
				CameraNum  int    `yaml:"CameraNum" json:"camera_num,omitempty"`
				CameraName string `yaml:"CameraName" json:"camera_name,omitempty"`
			} `yaml:"Cameras" json:"cameras,omitempty"`
			IsScenic bool `yaml:"IsScenic,omitempty" json:"is_scenic,omitempty"`
		} `yaml:"Groups" json:"groups,omitempty"`
	} `yaml:"CameraInfo" json:"camera_info"`
	RadioInfo struct {
		SelectedRadioNum int `yaml:"SelectedRadioNum" json:"selected_radio_num,omitempty"`
		Radios           []struct {
			RadioNum            int `yaml:"RadioNum" json:"radio_num,omitempty"`
			HopCount            int `yaml:"HopCount" json:"hop_count,omitempty"`
			NumFrequencies      int `yaml:"NumFrequencies" json:"num_frequencies,omitempty"`
			TunedToFrequencyNum int `yaml:"TunedToFrequencyNum" json:"tuned_to_frequency_num,omitempty"`
			ScanningIsOn        int `yaml:"ScanningIsOn" json:"scanning_is_on,omitempty"`
			Frequencies         []struct {
				FrequencyNum  int    `yaml:"FrequencyNum" json:"frequency_num,omitempty"`
				FrequencyName string `yaml:"FrequencyName" json:"frequency_name,omitempty"`
				Priority      int    `yaml:"Priority" json:"priority,omitempty"`
				CarIdx        int    `yaml:"CarIdx" json:"car_idx,omitempty"`
				EntryIdx      int    `yaml:"EntryIdx" json:"entry_idx,omitempty"`
				ClubID        int    `yaml:"ClubID" json:"club_id,omitempty"`
				CanScan       int    `yaml:"CanScan" json:"can_scan,omitempty"`
				CanSquawk     int    `yaml:"CanSquawk" json:"can_squawk,omitempty"`
				Muted         int    `yaml:"Muted" json:"muted,omitempty"`
				IsMutable     int    `yaml:"IsMutable" json:"is_mutable,omitempty"`
				IsDeletable   int    `yaml:"IsDeletable" json:"is_deletable,omitempty"`
			} `yaml:"Frequencies" json:"frequencies,omitempty"`
		} `yaml:"Radios" json:"radios,omitempty"`
	} `yaml:"RadioInfo" json:"radio_info"`
	DriverInfo struct {
		DriverCarIdx              int     `yaml:"DriverCarIdx" json:"driver_car_idx,omitempty"`
		DriverUserID              int     `yaml:"DriverUserID" json:"driver_user_id,omitempty"`
		PaceCarIdx                int     `yaml:"PaceCarIdx" json:"pace_car_idx,omitempty"`
		DriverHeadPosX            float64 `yaml:"DriverHeadPosX" json:"driver_head_pos_x,omitempty"`
		DriverHeadPosY            float64 `yaml:"DriverHeadPosY" json:"driver_head_pos_y,omitempty"`
		DriverHeadPosZ            float64 `yaml:"DriverHeadPosZ" json:"driver_head_pos_z,omitempty"`
		DriverCarIsElectric       int     `yaml:"DriverCarIsElectric" json:"driver_car_is_electric,omitempty"`
		DriverCarIdleRPM          float64 `yaml:"DriverCarIdleRPM" json:"driver_car_idle_rpm,omitempty"`
		DriverCarRedLine          float64 `yaml:"DriverCarRedLine" json:"driver_car_red_line,omitempty"`
		DriverCarEngCylinderCount int     `yaml:"DriverCarEngCylinderCount" json:"driver_car_eng_cylinder_count,omitempty"`
		DriverCarFuelKgPerLtr     float64 `yaml:"DriverCarFuelKgPerLtr" json:"driver_car_fuel_kg_per_ltr,omitempty"`
		DriverCarFuelMaxLtr       float64 `yaml:"DriverCarFuelMaxLtr" json:"driver_car_fuel_max_ltr,omitempty"`
		DriverCarMaxFuelPct       float64 `yaml:"DriverCarMaxFuelPct" json:"driver_car_max_fuel_pct,omitempty"`
		DriverCarGearNumForward   int     `yaml:"DriverCarGearNumForward" json:"driver_car_gear_num_forward,omitempty"`
		DriverCarGearNeutral      int     `yaml:"DriverCarGearNeutral" json:"driver_car_gear_neutral,omitempty"`
		DriverCarGearReverse      int     `yaml:"DriverCarGearReverse" json:"driver_car_gear_reverse,omitempty"`
		DriverCarSLFirstRPM       float64 `yaml:"DriverCarSLFirstRPM" json:"driver_car_sl_first_rpm,omitempty"`
		DriverCarSLShiftRPM       float64 `yaml:"DriverCarSLShiftRPM" json:"driver_car_sl_shift_rpm,omitempty"`
		DriverCarSLLastRPM        float64 `yaml:"DriverCarSLLastRPM" json:"driver_car_sl_last_rpm,omitempty"`
		DriverCarSLBlinkRPM       float64 `yaml:"DriverCarSLBlinkRPM" json:"driver_car_sl_blink_rpm,omitempty"`
		DriverCarVersion          string  `yaml:"DriverCarVersion" json:"driver_car_version,omitempty"`
		DriverPitTrkPct           float64 `yaml:"DriverPitTrkPct" json:"driver_pit_trk_pct,omitempty"`
		DriverCarEstLapTime       float64 `yaml:"DriverCarEstLapTime" json:"driver_car_est_lap_time,omitempty"`
		DriverSetupName           string  `yaml:"DriverSetupName" json:"driver_setup_name,omitempty"`
		DriverSetupIsModified     int     `yaml:"DriverSetupIsModified" json:"driver_setup_is_modified,omitempty"`
		DriverSetupLoadTypeName   string  `yaml:"DriverSetupLoadTypeName" json:"driver_setup_load_type_name,omitempty"`
		DriverSetupPassedTech     int     `yaml:"DriverSetupPassedTech" json:"driver_setup_passed_tech,omitempty"`
		DriverIncidentCount       int     `yaml:"DriverIncidentCount" json:"driver_incident_count,omitempty"`
		Drivers                   []struct {
			CarIdx                  int     `yaml:"CarIdx" json:"car_idx,omitempty"`
			UserName                string  `yaml:"UserName" json:"user_name,omitempty"`
			AbbrevName              string  `yaml:"AbbrevName" json:"abbrev_name,omitempty"`
			Initials                string  `yaml:"Initials" json:"initials,omitempty"`
			UserID                  int     `yaml:"UserID" json:"user_id,omitempty"`
			TeamID                  int     `yaml:"TeamID" json:"team_id,omitempty"`
			TeamName                string  `yaml:"TeamName" json:"team_name,omitempty"`
			CarNumber               string  `yaml:"CarNumber" json:"car_number,omitempty"`
			CarNumberRaw            int     `yaml:"CarNumberRaw" json:"car_number_raw,omitempty"`
			CarPath                 string  `yaml:"CarPath" json:"car_path,omitempty"`
			CarClassID              int     `yaml:"CarClassID" json:"car_class_id,omitempty"`
			CarID                   int     `yaml:"CarID" json:"car_id,omitempty"`
			CarIsPaceCar            int     `yaml:"CarIsPaceCar" json:"car_is_pace_car,omitempty"`
			CarIsAI                 int     `yaml:"CarIsAI" json:"car_is_ai,omitempty"`
			CarIsElectric           int     `yaml:"CarIsElectric" json:"car_is_electric,omitempty"`
			CarScreenName           string  `yaml:"CarScreenName" json:"car_screen_name,omitempty"`
			CarScreenNameShort      string  `yaml:"CarScreenNameShort" json:"car_screen_name_short,omitempty"`
			CarClassShortName       string  `yaml:"CarClassShortName" json:"car_class_short_name,omitempty"`
			CarClassRelSpeed        int     `yaml:"CarClassRelSpeed" json:"car_class_rel_speed,omitempty"`
			CarClassLicenseLevel    int     `yaml:"CarClassLicenseLevel" json:"car_class_license_level,omitempty"`
			CarClassMaxFuelPct      string  `yaml:"CarClassMaxFuelPct" json:"car_class_max_fuel_pct,omitempty"`
			CarClassWeightPenalty   string  `yaml:"CarClassWeightPenalty" json:"car_class_weight_penalty,omitempty"`
			CarClassPowerAdjust     string  `yaml:"CarClassPowerAdjust" json:"car_class_power_adjust,omitempty"`
			CarClassDryTireSetLimit string  `yaml:"CarClassDryTireSetLimit" json:"car_class_dry_tire_set_limit,omitempty"`
			CarClassColor           int     `yaml:"CarClassColor" json:"car_class_color,omitempty"`
			CarClassEstLapTime      float64 `yaml:"CarClassEstLapTime" json:"car_class_est_lap_time,omitempty"`
			IRating                 int     `yaml:"IRating" json:"i_rating,omitempty"`
			LicLevel                int     `yaml:"LicLevel" json:"lic_level,omitempty"`
			LicSubLevel             int     `yaml:"LicSubLevel" json:"lic_sub_level,omitempty"`
			LicString               string  `yaml:"LicString" json:"lic_string,omitempty"`
			LicColor                int     `yaml:"LicColor" json:"lic_color,omitempty"`
			IsSpectator             int     `yaml:"IsSpectator" json:"is_spectator,omitempty"`
			CarDesignStr            string  `yaml:"CarDesignStr" json:"car_design_str,omitempty"`
			HelmetDesignStr         string  `yaml:"HelmetDesignStr" json:"helmet_design_str,omitempty"`
			SuitDesignStr           string  `yaml:"SuitDesignStr" json:"suit_design_str,omitempty"`
			BodyType                int     `yaml:"BodyType" json:"body_type,omitempty"`
			FaceType                int     `yaml:"FaceType" json:"face_type,omitempty"`
			HelmetType              int     `yaml:"HelmetType" json:"helmet_type,omitempty"`
			CarNumberDesignStr      string  `yaml:"CarNumberDesignStr" json:"car_number_design_str,omitempty"`
			CarSponsor1             int     `yaml:"CarSponsor_1" json:"car_sponsor_1,omitempty"`
			CarSponsor2             int     `yaml:"CarSponsor_2" json:"car_sponsor_2,omitempty"`
			ClubName                string  `yaml:"ClubName" json:"club_name,omitempty"`
			ClubID                  int     `yaml:"ClubID" json:"club_id,omitempty"`
			DivisionName            string  `yaml:"DivisionName" json:"division_name,omitempty"`
			DivisionID              int     `yaml:"DivisionID" json:"division_id,omitempty"`
			CurDriverIncidentCount  int     `yaml:"CurDriverIncidentCount" json:"cur_driver_incident_count,omitempty"`
			TeamIncidentCount       int     `yaml:"TeamIncidentCount" json:"team_incident_count,omitempty"`
		} `yaml:"Drivers" json:"drivers,omitempty"`
	} `yaml:"DriverInfo" json:"driver_info"`
	SplitTimeInfo struct {
		Sectors []struct {
			SectorNum      int     `yaml:"SectorNum" json:"sector_num,omitempty"`
			SectorStartPct float64 `yaml:"SectorStartPct" json:"sector_start_pct,omitempty"`
		} `yaml:"Sectors" json:"sectors,omitempty"`
	} `yaml:"SplitTimeInfo" json:"split_time_info"`
	CarSetup struct {
		UpdateCount int `yaml:"UpdateCount" json:"update_count,omitempty"`
		TiresAero   struct {
			LeftFront struct {
				StartingPressure string `yaml:"StartingPressure" json:"starting_pressure,omitempty"`
				LastHotPressure  string `yaml:"LastHotPressure" json:"last_hot_pressure,omitempty"`
				LastTempsOMI     string `yaml:"LastTempsOMI" json:"last_temps_omi,omitempty"`
				TreadRemaining   string `yaml:"TreadRemaining" json:"tread_remaining,omitempty"`
			} `yaml:"LeftFront" json:"left_front"`
			LeftRear struct {
				StartingPressure string `yaml:"StartingPressure" json:"starting_pressure,omitempty"`
				LastHotPressure  string `yaml:"LastHotPressure" json:"last_hot_pressure,omitempty"`
				LastTempsOMI     string `yaml:"LastTempsOMI" json:"last_temps_omi,omitempty"`
				TreadRemaining   string `yaml:"TreadRemaining" json:"tread_remaining,omitempty"`
			} `yaml:"LeftRear" json:"left_rear"`
			RightFront struct {
				StartingPressure string `yaml:"StartingPressure" json:"starting_pressure,omitempty"`
				LastHotPressure  string `yaml:"LastHotPressure" json:"last_hot_pressure,omitempty"`
				LastTempsIMO     string `yaml:"LastTempsIMO" json:"last_temps_imo,omitempty"`
				TreadRemaining   string `yaml:"TreadRemaining" json:"tread_remaining,omitempty"`
			} `yaml:"RightFront" json:"right_front"`
			RightRear struct {
				StartingPressure string `yaml:"StartingPressure" json:"starting_pressure,omitempty"`
				LastHotPressure  string `yaml:"LastHotPressure" json:"last_hot_pressure,omitempty"`
				LastTempsIMO     string `yaml:"LastTempsIMO" json:"last_temps_imo,omitempty"`
				TreadRemaining   string `yaml:"TreadRemaining" json:"tread_remaining,omitempty"`
			} `yaml:"RightRear" json:"right_rear"`
			AeroBalanceCalc struct {
				FrontRhAtSpeed string `yaml:"FrontRhAtSpeed" json:"front_rh_at_speed,omitempty"`
				RearRhAtSpeed  string `yaml:"RearRhAtSpeed" json:"rear_rh_at_speed,omitempty"`
				WingSetting    string `yaml:"WingSetting" json:"wing_setting,omitempty"`
				FrontDownforce string `yaml:"FrontDownforce" json:"front_downforce,omitempty"`
			} `yaml:"AeroBalanceCalc" json:"aero_balance_calc"`
		} `yaml:"TiresAero" json:"tires_aero"`
		Chassis struct {
			Front struct {
				FarbSetting    int    `yaml:"FarbSetting" json:"farb_setting,omitempty"`
				TotalToeIn     string `yaml:"TotalToeIn" json:"total_toe_in,omitempty"`
				FuelLevel      string `yaml:"FuelLevel" json:"fuel_level,omitempty"`
				CrossWeight    string `yaml:"CrossWeight" json:"cross_weight,omitempty"`
				FrontMasterCyl string `yaml:"FrontMasterCyl" json:"front_master_cyl,omitempty"`
				RearMasterCyl  string `yaml:"RearMasterCyl" json:"rear_master_cyl,omitempty"`
				BrakePads      string `yaml:"BrakePads" json:"brake_pads,omitempty"`
			} `yaml:"Front" json:"front"`
			LeftFront struct {
				CornerWeight      string `yaml:"CornerWeight" json:"corner_weight,omitempty"`
				RideHeight        string `yaml:"RideHeight" json:"ride_height,omitempty"`
				SpringPerchOffset string `yaml:"SpringPerchOffset" json:"spring_perch_offset,omitempty"`
				SpringRate        string `yaml:"SpringRate" json:"spring_rate,omitempty"`
				Camber            string `yaml:"Camber" json:"camber,omitempty"`
			} `yaml:"LeftFront" json:"left_front"`
			LeftRear struct {
				CornerWeight      string `yaml:"CornerWeight" json:"corner_weight,omitempty"`
				RideHeight        string `yaml:"RideHeight" json:"ride_height,omitempty"`
				SpringPerchOffset string `yaml:"SpringPerchOffset" json:"spring_perch_offset,omitempty"`
				SpringRate        string `yaml:"SpringRate" json:"spring_rate,omitempty"`
				Camber            string `yaml:"Camber" json:"camber,omitempty"`
			} `yaml:"LeftRear" json:"left_rear"`
			Rear struct {
				RarbSetting int    `yaml:"RarbSetting" json:"rarb_setting,omitempty"`
				TotalToeIn  string `yaml:"TotalToeIn" json:"total_toe_in,omitempty"`
				WingSetting string `yaml:"WingSetting" json:"wing_setting,omitempty"`
			} `yaml:"Rear" json:"rear"`
			InCarDials struct {
				DisplayPage            string `yaml:"DisplayPage" json:"display_page,omitempty"`
				BrakePressureBias      string `yaml:"BrakePressureBias" json:"brake_pressure_bias,omitempty"`
				TractionControlSetting string `yaml:"TractionControlSetting" json:"traction_control_setting,omitempty"`
				AbsSetting             string `yaml:"AbsSetting" json:"abs_setting,omitempty"`
				ThrottleMapSetting     int    `yaml:"ThrottleMapSetting" json:"throttle_map_setting,omitempty"`
				NightLedStrips         string `yaml:"NightLedStrips" json:"night_led_strips,omitempty"`
			} `yaml:"InCarDials" json:"in_car_dials"`
			RightFront struct {
				CornerWeight      string `yaml:"CornerWeight" json:"corner_weight,omitempty"`
				RideHeight        string `yaml:"RideHeight" json:"ride_height,omitempty"`
				SpringPerchOffset string `yaml:"SpringPerchOffset" json:"spring_perch_offset,omitempty"`
				SpringRate        string `yaml:"SpringRate" json:"spring_rate,omitempty"`
				Camber            string `yaml:"Camber" json:"camber,omitempty"`
			} `yaml:"RightFront" json:"right_front"`
			RightRear struct {
				CornerWeight      string `yaml:"CornerWeight" json:"corner_weight,omitempty"`
				RideHeight        string `yaml:"RideHeight" json:"ride_height,omitempty"`
				SpringPerchOffset string `yaml:"SpringPerchOffset" json:"spring_perch_offset,omitempty"`
				SpringRate        string `yaml:"SpringRate" json:"spring_rate,omitempty"`
				Camber            string `yaml:"Camber" json:"camber,omitempty"`
			} `yaml:"RightRear" json:"right_rear"`
			GearsDifferential struct {
				GearStack     string `yaml:"GearStack" json:"gear_stack,omitempty"`
				FrictionFaces int    `yaml:"FrictionFaces" json:"friction_faces,omitempty"`
				DiffPreload   string `yaml:"DiffPreload" json:"diff_preload,omitempty"`
			} `yaml:"GearsDifferential" json:"gears_differential"`
		} `yaml:"Chassis" json:"chassis"`
		Dampers struct {
			FrontDampers struct {
				LowSpeedCompressionDamping  string `yaml:"LowSpeedCompressionDamping" json:"low_speed_compression_damping,omitempty"`
				HighSpeedCompressionDamping string `yaml:"HighSpeedCompressionDamping" json:"high_speed_compression_damping,omitempty"`
				LowSpeedReboundDamping      string `yaml:"LowSpeedReboundDamping" json:"low_speed_rebound_damping,omitempty"`
				HighSpeedReboundDamping     string `yaml:"HighSpeedReboundDamping" json:"high_speed_rebound_damping,omitempty"`
			} `yaml:"FrontDampers" json:"front_dampers"`
			RearDampers struct {
				LowSpeedCompressionDamping  string `yaml:"LowSpeedCompressionDamping" json:"low_speed_compression_damping,omitempty"`
				HighSpeedCompressionDamping string `yaml:"HighSpeedCompressionDamping" json:"high_speed_compression_damping,omitempty"`
				LowSpeedReboundDamping      string `yaml:"LowSpeedReboundDamping" json:"low_speed_rebound_damping,omitempty"`
				HighSpeedReboundDamping     string `yaml:"HighSpeedReboundDamping" json:"high_speed_rebound_damping,omitempty"`
			} `yaml:"RearDampers" json:"rear_dampers"`
		} `yaml:"Dampers" json:"dampers"`
	} `yaml:"CarSetup" json:"car_setup"`
}

type ResultsPositions struct {
	Position          int     `yaml:"Position" json:"position,omitempty"`
	ClassPosition     int     `yaml:"ClassPosition" json:"class_position,omitempty"`
	CarIdx            int     `yaml:"CarIdx" json:"car_idx,omitempty"`
	Lap               int     `yaml:"Lap" json:"lap,omitempty"`
	Time              float64 `yaml:"Time" json:"time,omitempty"`
	FastestLap        int     `yaml:"FastestLap" json:"fastest_lap,omitempty"`
	FastestTime       float64 `yaml:"FastestTime" json:"fastest_time,omitempty"`
	LastTime          float64 `yaml:"LastTime" json:"last_time,omitempty"`
	LapsLed           int     `yaml:"LapsLed" json:"laps_led,omitempty"`
	LapsComplete      int     `yaml:"LapsComplete" json:"laps_complete,omitempty"`
	JokerLapsComplete int     `yaml:"JokerLapsComplete" json:"joker_laps_complete,omitempty"`
	LapsDriven        float64 `yaml:"LapsDriven" json:"laps_driven,omitempty"`
	Incidents         int     `yaml:"Incidents" json:"incidents,omitempty"`
	ReasonOutID       int     `yaml:"ReasonOutId" json:"reason_out_id,omitempty"`
	ReasonOutStr      string  `yaml:"ReasonOutStr" json:"reason_out_str,omitempty"`
}
