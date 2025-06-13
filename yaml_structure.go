package irsdk

type Session struct {
	WeekendInfo        WeekendInfo        `yaml:"WeekendInfo"`
	SessionInfo        SessionInfo        `yaml:"SessionInfo"`
	QualifyResultsInfo QualifyResultsInfo `yaml:"QualifyResultsInfo"`
	CameraInfo         CameraInfo         `yaml:"CameraInfo"`
	RadioInfo          RadioInfo          `yaml:"RadioInfo"`
	DriverInfo         DriverInfo         `yaml:"DriverInfo"`
	SplitTimeInfo      SplitTimeInfo      `yaml:"SplitTimeInfo"`
	CarSetup           CarSetup           `yaml:"CarSetup"`
}
type Date struct {
}
type WeekendOptions struct {
	NumStarters                int    `yaml:"NumStarters"`
	StartingGrid               string `yaml:"StartingGrid"`
	QualifyScoring             string `yaml:"QualifyScoring"`
	CourseCautions             string `yaml:"CourseCautions"`
	StandingStart              int    `yaml:"StandingStart"`
	ShortParadeLap             int    `yaml:"ShortParadeLap"`
	Restarts                   string `yaml:"Restarts"`
	WeatherType                string `yaml:"WeatherType"`
	Skies                      string `yaml:"Skies"`
	WindDirection              string `yaml:"WindDirection"`
	WindSpeed                  string `yaml:"WindSpeed"`
	WeatherTemp                string `yaml:"WeatherTemp"`
	RelativeHumidity           string `yaml:"RelativeHumidity"`
	FogLevel                   string `yaml:"FogLevel"`
	TimeOfDay                  string `yaml:"TimeOfDay"`
	Date                       Date   `yaml:"Date"`
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
}
type TelemetryOptions struct {
	TelemetryDiskFile string `yaml:"TelemetryDiskFile"`
}
type WeekendInfo struct {
	TrackName              string           `yaml:"TrackName"`
	TrackID                int              `yaml:"TrackID"`
	TrackLength            string           `yaml:"TrackLength"`
	TrackLengthOfficial    string           `yaml:"TrackLengthOfficial"`
	TrackDisplayName       string           `yaml:"TrackDisplayName"`
	TrackDisplayShortName  string           `yaml:"TrackDisplayShortName"`
	TrackConfigName        string           `yaml:"TrackConfigName"`
	TrackCity              string           `yaml:"TrackCity"`
	TrackState             string           `yaml:"TrackState"`
	TrackCountry           string           `yaml:"TrackCountry"`
	TrackAltitude          string           `yaml:"TrackAltitude"`
	TrackLatitude          string           `yaml:"TrackLatitude"`
	TrackLongitude         string           `yaml:"TrackLongitude"`
	TrackNorthOffset       string           `yaml:"TrackNorthOffset"`
	TrackNumTurns          int              `yaml:"TrackNumTurns"`
	TrackPitSpeedLimit     string           `yaml:"TrackPitSpeedLimit"`
	TrackPaceSpeed         string           `yaml:"TrackPaceSpeed"`
	TrackNumPitStalls      int              `yaml:"TrackNumPitStalls"`
	TrackType              string           `yaml:"TrackType"`
	TrackDirection         string           `yaml:"TrackDirection"`
	TrackWeatherType       string           `yaml:"TrackWeatherType"`
	TrackSkies             string           `yaml:"TrackSkies"`
	TrackSurfaceTemp       string           `yaml:"TrackSurfaceTemp"`
	TrackSurfaceTempCrew   string           `yaml:"TrackSurfaceTempCrew"`
	TrackAirTemp           string           `yaml:"TrackAirTemp"`
	TrackAirPressure       string           `yaml:"TrackAirPressure"`
	TrackAirDensity        string           `yaml:"TrackAirDensity"`
	TrackWindVel           string           `yaml:"TrackWindVel"`
	TrackWindDir           string           `yaml:"TrackWindDir"`
	TrackRelativeHumidity  string           `yaml:"TrackRelativeHumidity"`
	TrackFogLevel          string           `yaml:"TrackFogLevel"`
	TrackPrecipitation     string           `yaml:"TrackPrecipitation"`
	TrackCleanup           int              `yaml:"TrackCleanup"`
	TrackDynamicTrack      int              `yaml:"TrackDynamicTrack"`
	TrackVersion           string           `yaml:"TrackVersion"`
	SeriesID               int              `yaml:"SeriesID"`
	SeasonID               int              `yaml:"SeasonID"`
	SessionID              int              `yaml:"SessionID"`
	SubSessionID           int              `yaml:"SubSessionID"`
	LeagueID               int              `yaml:"LeagueID"`
	Official               int              `yaml:"Official"`
	RaceWeek               int              `yaml:"RaceWeek"`
	EventType              string           `yaml:"EventType"`
	Category               string           `yaml:"Category"`
	SimMode                string           `yaml:"SimMode"`
	TeamRacing             int              `yaml:"TeamRacing"`
	MinDrivers             int              `yaml:"MinDrivers"`
	MaxDrivers             int              `yaml:"MaxDrivers"`
	DCRuleSet              string           `yaml:"DCRuleSet"`
	QualifierMustStartRace int              `yaml:"QualifierMustStartRace"`
	NumCarClasses          int              `yaml:"NumCarClasses"`
	NumCarTypes            int              `yaml:"NumCarTypes"`
	HeatRacing             int              `yaml:"HeatRacing"`
	BuildType              string           `yaml:"BuildType"`
	BuildTarget            string           `yaml:"BuildTarget"`
	BuildVersion           string           `yaml:"BuildVersion"`
	RaceFarm               string           `yaml:"RaceFarm"`
	WeekendOptions         WeekendOptions   `yaml:"WeekendOptions"`
	TelemetryOptions       TelemetryOptions `yaml:"TelemetryOptions"`
}
type ResultsPositions struct {
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
}
type ResultsFastestLap struct {
	CarIdx      int     `yaml:"CarIdx"`
	FastestLap  int     `yaml:"FastestLap"`
	FastestTime float64 `yaml:"FastestTime"`
}
type Sessions struct {
	SessionNum                       int                 `yaml:"SessionNum"`
	SessionLaps                      string              `yaml:"SessionLaps"`
	SessionTime                      string              `yaml:"SessionTime"`
	SessionNumLapsToAvg              int                 `yaml:"SessionNumLapsToAvg"`
	SessionType                      string              `yaml:"SessionType"`
	SessionTrackRubberState          string              `yaml:"SessionTrackRubberState"`
	SessionName                      string              `yaml:"SessionName"`
	SessionSubType                   interface{}         `yaml:"SessionSubType"`
	SessionSkipped                   int                 `yaml:"SessionSkipped"`
	SessionRunGroupsUsed             int                 `yaml:"SessionRunGroupsUsed"`
	SessionEnforceTireCompoundChange int                 `yaml:"SessionEnforceTireCompoundChange"`
	ResultsPositions                 []ResultsPositions  `yaml:"ResultsPositions"`
	ResultsFastestLap                []ResultsFastestLap `yaml:"ResultsFastestLap"`
	ResultsAverageLapTime            int                 `yaml:"ResultsAverageLapTime"`
	ResultsNumCautionFlags           int                 `yaml:"ResultsNumCautionFlags"`
	ResultsNumCautionLaps            int                 `yaml:"ResultsNumCautionLaps"`
	ResultsNumLeadChanges            int                 `yaml:"ResultsNumLeadChanges"`
	ResultsLapsComplete              int                 `yaml:"ResultsLapsComplete"`
	ResultsOfficial                  int                 `yaml:"ResultsOfficial"`
}
type SessionInfo struct {
	CurrentSessionNum int        `yaml:"CurrentSessionNum"`
	Sessions          []Sessions `yaml:"Sessions"`
}
type Results struct {
	Position      int     `yaml:"Position"`
	ClassPosition int     `yaml:"ClassPosition"`
	CarIdx        int     `yaml:"CarIdx"`
	FastestLap    int     `yaml:"FastestLap"`
	FastestTime   float64 `yaml:"FastestTime"`
}
type QualifyResultsInfo struct {
	Results []Results `yaml:"Results"`
}
type Cameras struct {
	CameraNum  int    `yaml:"CameraNum"`
	CameraName string `yaml:"CameraName"`
}
type Groups struct {
	GroupNum  int       `yaml:"GroupNum"`
	GroupName string    `yaml:"GroupName"`
	Cameras   []Cameras `yaml:"Cameras"`
	IsScenic  bool      `yaml:"IsScenic,omitempty"`
}
type CameraInfo struct {
	Groups []Groups `yaml:"Groups"`
}
type Frequencies struct {
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
}
type Radios struct {
	RadioNum            int           `yaml:"RadioNum"`
	HopCount            int           `yaml:"HopCount"`
	NumFrequencies      int           `yaml:"NumFrequencies"`
	TunedToFrequencyNum int           `yaml:"TunedToFrequencyNum"`
	ScanningIsOn        int           `yaml:"ScanningIsOn"`
	Frequencies         []Frequencies `yaml:"Frequencies"`
}
type RadioInfo struct {
	SelectedRadioNum int      `yaml:"SelectedRadioNum"`
	Radios           []Radios `yaml:"Radios"`
}
type DriverTires struct {
	TireIndex        int    `yaml:"TireIndex"`
	TireCompoundType string `yaml:"TireCompoundType"`
}
type Drivers struct {
	CarIdx                  int         `yaml:"CarIdx"`
	UserName                string      `yaml:"UserName"`
	AbbrevName              interface{} `yaml:"AbbrevName"`
	Initials                interface{} `yaml:"Initials"`
	UserID                  int         `yaml:"UserID"`
	TeamID                  int         `yaml:"TeamID"`
	TeamName                string      `yaml:"TeamName"`
	CarNumber               string      `yaml:"CarNumber"`
	CarNumberRaw            int         `yaml:"CarNumberRaw"`
	CarPath                 string      `yaml:"CarPath"`
	CarClassID              int         `yaml:"CarClassID"`
	CarID                   int         `yaml:"CarID"`
	CarIsPaceCar            int         `yaml:"CarIsPaceCar"`
	CarIsAI                 int         `yaml:"CarIsAI"`
	CarIsElectric           int         `yaml:"CarIsElectric"`
	CarScreenName           string      `yaml:"CarScreenName"`
	CarScreenNameShort      string      `yaml:"CarScreenNameShort"`
	CarCfg                  int         `yaml:"CarCfg"`
	CarCfgName              interface{} `yaml:"CarCfgName"`
	CarCfgCustomPaintExt    interface{} `yaml:"CarCfgCustomPaintExt"`
	CarClassShortName       interface{} `yaml:"CarClassShortName"`
	CarClassRelSpeed        int         `yaml:"CarClassRelSpeed"`
	CarClassLicenseLevel    int         `yaml:"CarClassLicenseLevel"`
	CarClassMaxFuelPct      string      `yaml:"CarClassMaxFuelPct"`
	CarClassWeightPenalty   string      `yaml:"CarClassWeightPenalty"`
	CarClassPowerAdjust     string      `yaml:"CarClassPowerAdjust"`
	CarClassDryTireSetLimit string      `yaml:"CarClassDryTireSetLimit"`
	CarClassColor           int         `yaml:"CarClassColor"`
	CarClassEstLapTime      float64     `yaml:"CarClassEstLapTime"`
	IRating                 int         `yaml:"IRating"`
	LicLevel                int         `yaml:"LicLevel"`
	LicSubLevel             int         `yaml:"LicSubLevel"`
	LicString               string      `yaml:"LicString"`
	LicColor                int         `yaml:"LicColor"`
	IsSpectator             int         `yaml:"IsSpectator"`
	CarDesignStr            string      `yaml:"CarDesignStr"`
	HelmetDesignStr         string      `yaml:"HelmetDesignStr"`
	SuitDesignStr           string      `yaml:"SuitDesignStr"`
	BodyType                int         `yaml:"BodyType"`
	FaceType                int         `yaml:"FaceType"`
	HelmetType              int         `yaml:"HelmetType"`
	CarNumberDesignStr      string      `yaml:"CarNumberDesignStr"`
	CarSponsor1             int         `yaml:"CarSponsor_1"`
	CarSponsor2             int         `yaml:"CarSponsor_2"`
	ClubName                string      `yaml:"ClubName"`
	ClubID                  int         `yaml:"ClubID"`
	FlairName               string      `yaml:"FlairName"`
	FlairID                 int         `yaml:"FlairID"`
	DivisionName            string      `yaml:"DivisionName"`
	DivisionID              int         `yaml:"DivisionID"`
	CurDriverIncidentCount  int         `yaml:"CurDriverIncidentCount"`
	TeamIncidentCount       int         `yaml:"TeamIncidentCount"`
}
type DriverInfo struct {
	DriverCarIdx              int           `yaml:"DriverCarIdx"`
	DriverUserID              int           `yaml:"DriverUserID"`
	PaceCarIdx                int           `yaml:"PaceCarIdx"`
	DriverHeadPosX            float64       `yaml:"DriverHeadPosX"`
	DriverHeadPosY            float64       `yaml:"DriverHeadPosY"`
	DriverHeadPosZ            float64       `yaml:"DriverHeadPosZ"`
	DriverCarIsElectric       int           `yaml:"DriverCarIsElectric"`
	DriverCarIdleRPM          float64       `yaml:"DriverCarIdleRPM"`
	DriverCarRedLine          float64       `yaml:"DriverCarRedLine"`
	DriverCarEngCylinderCount int           `yaml:"DriverCarEngCylinderCount"`
	DriverCarFuelKgPerLtr     float64       `yaml:"DriverCarFuelKgPerLtr"`
	DriverCarFuelMaxLtr       float64       `yaml:"DriverCarFuelMaxLtr"`
	DriverCarMaxFuelPct       float64       `yaml:"DriverCarMaxFuelPct"`
	DriverCarGearNumForward   int           `yaml:"DriverCarGearNumForward"`
	DriverCarGearNeutral      int           `yaml:"DriverCarGearNeutral"`
	DriverCarGearReverse      int           `yaml:"DriverCarGearReverse"`
	DriverCarSLFirstRPM       float64       `yaml:"DriverCarSLFirstRPM"`
	DriverCarSLShiftRPM       float64       `yaml:"DriverCarSLShiftRPM"`
	DriverCarSLLastRPM        float64       `yaml:"DriverCarSLLastRPM"`
	DriverCarSLBlinkRPM       float64       `yaml:"DriverCarSLBlinkRPM"`
	DriverCarVersion          string        `yaml:"DriverCarVersion"`
	DriverPitTrkPct           float64       `yaml:"DriverPitTrkPct"`
	DriverCarEstLapTime       float64       `yaml:"DriverCarEstLapTime"`
	DriverSetupName           string        `yaml:"DriverSetupName"`
	DriverSetupIsModified     int           `yaml:"DriverSetupIsModified"`
	DriverSetupLoadTypeName   string        `yaml:"DriverSetupLoadTypeName"`
	DriverSetupPassedTech     int           `yaml:"DriverSetupPassedTech"`
	DriverIncidentCount       int           `yaml:"DriverIncidentCount"`
	DriverBrakeCurvingFactor  float64       `yaml:"DriverBrakeCurvingFactor"`
	DriverTires               []DriverTires `yaml:"DriverTires"`
	Drivers                   []Drivers     `yaml:"Drivers"`
}
type Sectors struct {
	SectorNum      int     `yaml:"SectorNum"`
	SectorStartPct float64 `yaml:"SectorStartPct"`
}
type SplitTimeInfo struct {
	Sectors []Sectors `yaml:"Sectors"`
}
type TireLeftFront struct {
	ColdPressure    string `yaml:"ColdPressure"`
	LastHotPressure string `yaml:"LastHotPressure"`
	LastTempsOMI    string `yaml:"LastTempsOMI"`
	TreadRemaining  string `yaml:"TreadRemaining"`
}
type TireLeftRear struct {
	ColdPressure    string `yaml:"ColdPressure"`
	LastHotPressure string `yaml:"LastHotPressure"`
	LastTempsOMI    string `yaml:"LastTempsOMI"`
	TreadRemaining  string `yaml:"TreadRemaining"`
}
type TireRightFront struct {
	ColdPressure    string `yaml:"ColdPressure"`
	LastHotPressure string `yaml:"LastHotPressure"`
	LastTempsIMO    string `yaml:"LastTempsIMO"`
	TreadRemaining  string `yaml:"TreadRemaining"`
}
type TireRightRear struct {
	ColdPressure    string `yaml:"ColdPressure"`
	LastHotPressure string `yaml:"LastHotPressure"`
	LastTempsIMO    string `yaml:"LastTempsIMO"`
	TreadRemaining  string `yaml:"TreadRemaining"`
}
type Tires struct {
	LeftFront  TireLeftFront  `yaml:"LeftFront"`
	LeftRear   TireLeftRear   `yaml:"LeftRear"`
	RightFront TireRightFront `yaml:"RightFront"`
	RightRear  TireRightRear  `yaml:"RightRear"`
}
type Front struct {
	BallastForward    string `yaml:"BallastForward"`
	NoseWeight        string `yaml:"NoseWeight"`
	CrossWeight       string `yaml:"CrossWeight"`
	SteeringRatio     int    `yaml:"SteeringRatio"`
	SteeringOffset    string `yaml:"SteeringOffset"`
	FrontBrakeBias    string `yaml:"FrontBrakeBias"`
	TapeConfiguration string `yaml:"TapeConfiguration"`
}
type LeftFront struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	ShockDeflection   string `yaml:"ShockDeflection"`
	SpringDeflection  string `yaml:"SpringDeflection"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
	Caster            string `yaml:"Caster"`
	ToeIn             string `yaml:"ToeIn"`
}
type LeftRear struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	ShockDeflection   string `yaml:"ShockDeflection"`
	SpringDeflection  string `yaml:"SpringDeflection"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	LeftRearToeIn     string `yaml:"LeftRearToeIn"`
	Camber            string `yaml:"Camber"`
	TrackBarHeight    string `yaml:"TrackBarHeight"`
	TruckArmMount     string `yaml:"TruckArmMount"`
}
type FrontArb struct {
	Diameter         string `yaml:"Diameter"`
	ArmAsymmetry     string `yaml:"ArmAsymmetry"`
	ChainOrSolidLink string `yaml:"ChainOrSolidLink"`
	LinkSlack        string `yaml:"LinkSlack"`
	Preload          string `yaml:"Preload"`
	Attach           int    `yaml:"Attach"`
}
type RightFront struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	ShockDeflection   string `yaml:"ShockDeflection"`
	SpringDeflection  string `yaml:"SpringDeflection"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
	Caster            string `yaml:"Caster"`
	ToeIn             string `yaml:"ToeIn"`
}
type RightRear struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	ShockDeflection   string `yaml:"ShockDeflection"`
	SpringDeflection  string `yaml:"SpringDeflection"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	RightRearToeIn    string `yaml:"RightRearToeIn"`
	Camber            string `yaml:"Camber"`
	TrackBarHeight    string `yaml:"TrackBarHeight"`
	TruckArmMount     string `yaml:"TruckArmMount"`
	TruckArmPreload   string `yaml:"TruckArmPreload"`
}
type Rear struct {
	RearEndRatio      float64 `yaml:"RearEndRatio"`
	ArbDiameter       string  `yaml:"ArbDiameter"`
	ArmAsymmetry      string  `yaml:"ArmAsymmetry"`
	SliderOrSolidLink string  `yaml:"SliderOrSolidLink"`
	Preload           string  `yaml:"Preload"`
	LinkSlack         string  `yaml:"LinkSlack"`
}
type Chassis struct {
	Front      Front      `yaml:"Front"`
	LeftFront  LeftFront  `yaml:"LeftFront"`
	LeftRear   LeftRear   `yaml:"LeftRear"`
	FrontArb   FrontArb   `yaml:"FrontArb"`
	RightFront RightFront `yaml:"RightFront"`
	RightRear  RightRear  `yaml:"RightRear"`
	Rear       Rear       `yaml:"Rear"`
}
type CarSetup struct {
	UpdateCount int     `yaml:"UpdateCount"`
	Tires       Tires   `yaml:"Tires"`
	Chassis     Chassis `yaml:"Chassis"`
}
