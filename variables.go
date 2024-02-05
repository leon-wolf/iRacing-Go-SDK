package irsdk

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type varBuffer struct {
	tickCount int // used to detect changes in data
	bufOffset int // offset from header
}

type Variable struct {
	varType     int // irsdk_VarType
	offset      int // offset fron start of buffer row
	count       int // number of entrys (array) so length in bytes would be irsdk_VarTypeBytes[type] * count
	countAsTime bool
	Name        string
	Desc        string
	Unit        string
	Value       interface{}
	rawBytes    []byte
}

func (v Variable) String() string {
	var ret string
	switch v.varType {
	case 0:
		ret = fmt.Sprintf("%c", v.Value)
	case 1:
		ret = fmt.Sprintf("%v", v.Value)
	case 2:
		ret = fmt.Sprintf("%d", v.Value)
	case 3:
		ret = fmt.Sprintf("%s", v.Value)
	case 4:
		ret = fmt.Sprintf("%f", v.Value)
	case 5:
		ret = fmt.Sprintf("%f", v.Value)
	default:
		ret = fmt.Sprintf("Unknown (%d)", v.varType)
	}
	return ret
}

// TelemetryVars holds all variables we can read from telemetry live
type TelemetryVars struct {
	lastVersion int
	vars        map[string]Variable
	mux         sync.Mutex
}

func findLatestBuffer(r reader, h *header) varBuffer {
	var vb varBuffer
	foundTickCount := 0
	for i := 0; i < h.numBuf; i++ {
		rbuf := make([]byte, 16)
		_, err := r.ReadAt(rbuf, int64(48+i*16))
		if err != nil {
			log.Fatal(err)
		}
		currentVb := varBuffer{
			byte4ToInt(rbuf[0:4]),
			byte4ToInt(rbuf[4:8]),
		}
		//fmt.Printf("BUFF?: %+v\n", currentVb)
		if foundTickCount < currentVb.tickCount {
			foundTickCount = currentVb.tickCount
			vb = currentVb
		}
	}
	//fmt.Printf("BUFF: %+v\n", vb)
	return vb
}

func readVariableHeaders(r reader, h *header) *TelemetryVars {
	vars := TelemetryVars{vars: make(map[string]Variable, h.numVars)}
	for i := 0; i < h.numVars; i++ {
		rbuf := make([]byte, 144)
		_, err := r.ReadAt(rbuf, int64(h.headerOffset+i*144))
		if err != nil {
			log.Fatal(err)
		}
		v := Variable{
			byte4ToInt(rbuf[0:4]),
			byte4ToInt(rbuf[4:8]),
			byte4ToInt(rbuf[8:12]),
			int(rbuf[12]) > 0,
			bytesToString(rbuf[16:48]),
			bytesToString(rbuf[48:112]),
			bytesToString(rbuf[112:144]),
			nil,
			nil,
		}
		vars.vars[v.Name] = v
	}
	return &vars
}

func readVariableValues(sdk *IRSDK) bool {
	newData := false
	if sessionStatusOK(sdk.h.status) {
		// find latest buffer for variables
		vb := findLatestBuffer(sdk.r, sdk.h)
		sdk.tVars.mux.Lock()
		if sdk.tVars.lastVersion < vb.tickCount {
			newData = true
			sdk.tVars.lastVersion = vb.tickCount
			sdk.lastValidData = time.Now().Unix()
			for varName, v := range sdk.tVars.vars {
				var rbuf []byte
				switch v.varType {
				case 0:
					rbuf = make([]byte, 1)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = string(rbuf[0])
				case 1:
					rbuf = make([]byte, 1)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = int(rbuf[0]) > 0
				case 2:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = byte4ToInt(rbuf)
				case 3:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = byte4toBitField(rbuf)
				case 4:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = byte4ToFloat(rbuf)
				case 5:
					rbuf = make([]byte, 8)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = byte8ToFloat(rbuf)
				}
				v.rawBytes = rbuf
				sdk.tVars.vars[varName] = v
			}
		}
		sdk.tVars.mux.Unlock()
	}

	return newData
}

const AirDensity = "AirDensity"                                           // Density of air at start/finish line, kg/m^3
const AirPressure = "AirPressure"                                         // Pressure of air at start/finish line, Pa
const AirTemp = "AirTemp"                                                 // Temperature of air at start/finish line, C
const Brake = "Brake"                                                     // 0=brake released to 1=max pedal force, %
const BrakeABSactive = "BrakeABSactive"                                   // true if abs is currently reducing brake force pressure,
const BrakeRaw = "BrakeRaw"                                               // Raw brake input 0=brake released to 1=max pedal force, %
const CamCameraNumber = "CamCameraNumber"                                 // Active camera number,
const CamCameraState = "CamCameraState"                                   // State of camera system, irsdk_CameraState
const CamCarIdx = "CamCarIdx"                                             // Active camera's focus car index,
const CamGroupNumber = "CamGroupNumber"                                   // Active camera group number,
const CarIdxBestLapNum = "CarIdxBestLapNum"                               // Cars best lap number,
const CarIdxBestLapTime = "CarIdxBestLapTime"                             // Cars best lap time, s
const CarIdxClass = "CarIdxClass"                                         // Cars class id by car index,
const CarIdxClassPosition = "CarIdxClassPosition"                         // Cars class position in race by car index,
const CarIdxEstTime = "CarIdxEstTime"                                     // Estimated time to reach current location on track, s
const CarIdxF2Time = "CarIdxF2Time"                                       // Race time behind leader or fastest lap time otherwise, s
const CarIdxFastRepairsUsed = "CarIdxFastRepairsUsed"                     // How many fast repairs each car has used,
const CarIdxGear = "CarIdxGear"                                           // -1=reverse  0=neutral  1..n=current gear by car index,
const CarIdxLap = "CarIdxLap"                                             // Laps started by car index,
const CarIdxLapCompleted = "CarIdxLapCompleted"                           // Laps completed by car index,
const CarIdxLapDistPct = "CarIdxLapDistPct"                               // Percentage distance around lap by car index, %
const CarIdxLastLapTime = "CarIdxLastLapTime"                             // Cars last lap time, s
const CarIdxOnPitRoad = "CarIdxOnPitRoad"                                 // On pit road between the cones by car index,
const CarIdxP2P_Count = "CarIdxP2P_Count"                                 // Push2Pass count of usage (or remaining in Race),
const CarIdxP2P_Status = "CarIdxP2P_Status"                               // Push2Pass active or not,
const CarIdxPaceFlags = "CarIdxPaceFlags"                                 // Pacing status flags for each car, irsdk_PaceFlags
const CarIdxPaceLine = "CarIdxPaceLine"                                   // What line cars are pacing in  or -1 if not pacing,
const CarIdxPaceRow = "CarIdxPaceRow"                                     // What row cars are pacing in  or -1 if not pacing,
const CarIdxPosition = "CarIdxPosition"                                   // Cars position in race by car index,
const CarIdxQualTireCompound = "CarIdxQualTireCompound"                   // Cars Qual tire compound,
const CarIdxQualTireCompoundLocked = "CarIdxQualTireCompoundLocked"       // Cars Qual tire compound is locked-in,
const CarIdxRPM = "CarIdxRPM"                                             // Engine rpm by car index, revs/min
const CarIdxSessionFlags = "CarIdxSessionFlags"                           // Session flags for each player, irsdk_Flags
const CarIdxSteer = "CarIdxSteer"                                         // Steering wheel angle by car index, rad
const CarIdxTireCompound = "CarIdxTireCompound"                           // Cars current tire compound,
const CarIdxTrackSurface = "CarIdxTrackSurface"                           // Track surface type by car index, irsdk_TrkLoc
const CarIdxTrackSurfaceMaterial = "CarIdxTrackSurfaceMaterial"           // Track surface material type by car index, irsdk_TrkSurf
const CarLeftRight = "CarLeftRight"                                       // Notify if car is to the left or right of driver, irsdk_CarLeftRight
const ChanAvgLatency = "ChanAvgLatency"                                   // Communications average latency, s
const ChanClockSkew = "ChanClockSkew"                                     // Communications server clock skew, s
const ChanLatency = "ChanLatency"                                         // Communications latency, s
const ChanPartnerQuality = "ChanPartnerQuality"                           // Partner communications quality, %
const ChanQuality = "ChanQuality"                                         // Communications quality, %
const Clutch = "Clutch"                                                   // 0=disengaged to 1=fully engaged, %
const ClutchRaw = "ClutchRaw"                                             // Raw clutch input 0=disengaged to 1=fully engaged, %
const CpuUsageBG = "CpuUsageBG"                                           // Percent of available tim bg thread took with a 1 sec avg, %
const CpuUsageFG = "CpuUsageFG"                                           // Percent of available tim fg thread took with a 1 sec avg, %
const DCDriversSoFar = "DCDriversSoFar"                                   // Number of team drivers who have run a stint,
const DCLapStatus = "DCLapStatus"                                         // Status of driver change lap requirements,
const dcPitSpeedLimiterToggle = "dcPitSpeedLimiterToggle"                 // In car traction control active,
const dcStarter = "dcStarter"                                             // In car trigger car starter,
const DisplayUnits = "DisplayUnits"                                       // Default units for the user interface 0 = english 1 = metric,
const dpFastRepair = "dpFastRepair"                                       // Pitstop fast repair set,
const dpFuelAddKg = "dpFuelAddKg"                                         // Pitstop fuel add amount, kg
const dpFuelAutoFillActive = "dpFuelAutoFillActive"                       // Pitstop auto fill fuel next stop flag,
const dpFuelAutoFillEnabled = "dpFuelAutoFillEnabled"                     // Pitstop auto fill fuel system enabled,
const dpFuelFill = "dpFuelFill"                                           // Pitstop fuel fill flag,
const dpLFTireChange = "dpLFTireChange"                                   // Pitstop lf tire change request,
const dpLFTireColdPress = "dpLFTireColdPress"                             // Pitstop lf tire cold pressure adjustment, Pa
const dpLRTireChange = "dpLRTireChange"                                   // Pitstop lr tire change request,
const dpLRTireColdPress = "dpLRTireColdPress"                             // Pitstop lr tire cold pressure adjustment, Pa
const dpRFTireChange = "dpRFTireChange"                                   // Pitstop rf tire change request,
const dpRFTireColdPress = "dpRFTireColdPress"                             // Pitstop rf cold tire pressure adjustment, Pa
const dpRRTireChange = "dpRRTireChange"                                   // Pitstop rr tire change request,
const dpRRTireColdPress = "dpRRTireColdPress"                             // Pitstop rr cold tire pressure adjustment, Pa
const dpWindshieldTearoff = "dpWindshieldTearoff"                         // Pitstop windshield tearoff,
const DriverMarker = "DriverMarker"                                       // Driver activated flag,
const Engine0_RPM = "Engine0_RPM"                                         // Engine0Engine rpm, revs/min
const EngineWarnings = "EngineWarnings"                                   // Bitfield for warning lights, irsdk_EngineWarnings
const EnterExitReset = "EnterExitReset"                                   // Indicate action the reset key will take 0 enter 1 exit 2 reset,
const FastRepairAvailable = "FastRepairAvailable"                         // How many fast repairs left  255 is unlimited,
const FastRepairUsed = "FastRepairUsed"                                   // How many fast repairs used so far,
const FogLevel = "FogLevel"                                               // Fog level at start/finish line, %
const FrameRate = "FrameRate"                                             // Average frames per second, fps
const FrontTireSetsAvailable = "FrontTireSetsAvailable"                   // How many front tire sets are remaining  255 is unlimited,
const FrontTireSetsUsed = "FrontTireSetsUsed"                             // How many front tire sets used so far,
const FuelLevel = "FuelLevel"                                             // Liters of fuel remaining, l
const FuelLevelPct = "FuelLevelPct"                                       // Percent fuel remaining, %
const FuelPress = "FuelPress"                                             // Engine fuel pressure, bar
const FuelUsePerHour = "FuelUsePerHour"                                   // Engine fuel used instantaneous, kg/h
const Gear = "Gear"                                                       // -1=reverse  0=neutral  1..n=current gear,
const GpuUsage = "GpuUsage"                                               // Percent of available tim gpu took with a 1 sec avg, %
const HandbrakeRaw = "HandbrakeRaw"                                       // Raw handbrake input 0=handbrake released to 1=max force, %
const IsDiskLoggingActive = "IsDiskLoggingActive"                         // 0=disk based telemetry file not being written  1=being written,
const IsDiskLoggingEnabled = "IsDiskLoggingEnabled"                       // 0=disk based telemetry turned off  1=turned on,
const IsGarageVisible = "IsGarageVisible"                                 // 1=Garage screen is visible,
const IsInGarage = "IsInGarage"                                           // 1=Car in garage physics running,
const IsOnTrack = "IsOnTrack"                                             // 1=Car on track physics running with player in car,
const IsOnTrackCar = "IsOnTrackCar"                                       // 1=Car on track physics running,
const IsReplayPlaying = "IsReplayPlaying"                                 // 0=replay not playing  1=replay playing,
const Lap = "Lap"                                                         // Laps started count,
const LapBestLap = "LapBestLap"                                           // Players best lap number,
const LapBestLapTime = "LapBestLapTime"                                   // Players best lap time, s
const LapBestNLapLap = "LapBestNLapLap"                                   // Player last lap in best N average lap time,
const LapBestNLapTime = "LapBestNLapTime"                                 // Player best N average lap time, s
const LapCompleted = "LapCompleted"                                       // Laps completed count,
const LapCurrentLapTime = "LapCurrentLapTime"                             // Estimate of players current lap time as shown in F3 box, s
const LapDeltaToBestLap = "LapDeltaToBestLap"                             // Delta time for best lap, s
const LapDeltaToBestLap_DD = "LapDeltaToBestLap_DD"                       // Rate of change of delta time for best lap, s/s
const LapDeltaToBestLap_OK = "LapDeltaToBestLap_OK"                       // Delta time for best lap is valid,
const LapDeltaToOptimalLap = "LapDeltaToOptimalLap"                       // Delta time for optimal lap, s
const LapDeltaToOptimalLap_DD = "LapDeltaToOptimalLap_DD"                 // Rate of change of delta time for optimal lap, s/s
const LapDeltaToOptimalLap_OK = "LapDeltaToOptimalLap_OK"                 // Delta time for optimal lap is valid,
const LapDeltaToSessionBestLap = "LapDeltaToSessionBestLap"               // Delta time for session best lap, s
const LapDeltaToSessionBestLap_DD = "LapDeltaToSessionBestLap_DD"         // Rate of change of delta time for session best lap, s/s
const LapDeltaToSessionBestLap_OK = "LapDeltaToSessionBestLap_OK"         // Delta time for session best lap is valid,
const LapDeltaToSessionLastlLap = "LapDeltaToSessionLastlLap"             // Delta time for session last lap, s
const LapDeltaToSessionLastlLap_DD = "LapDeltaToSessionLastlLap_DD"       // Rate of change of delta time for session last lap, s/s
const LapDeltaToSessionLastlLap_OK = "LapDeltaToSessionLastlLap_OK"       // Delta time for session last lap is valid,
const LapDeltaToSessionOptimalLap = "LapDeltaToSessionOptimalLap"         // Delta time for session optimal lap, s
const LapDeltaToSessionOptimalLap_DD = "LapDeltaToSessionOptimalLap_DD"   // Rate of change of delta time for session optimal lap, s/s
const LapDeltaToSessionOptimalLap_OK = "LapDeltaToSessionOptimalLap_OK"   // Delta time for session optimal lap is valid,
const LapDist = "LapDist"                                                 // Meters traveled from S/F this lap, m
const LapDistPct = "LapDistPct"                                           // Percentage distance around lap, %
const LapLasNLapSeq = "LapLasNLapSeq"                                     // Player num consecutive clean laps completed for N average,
const LapLastLapTime = "LapLastLapTime"                                   // Players last lap time, s
const LapLastNLapTime = "LapLastNLapTime"                                 // Player last N average lap time, s
const LatAccel = "LatAccel"                                               // Lateral acceleration (including gravity), m/s^2
const LatAccel_ST = "LatAccel_ST"                                         // Lateral acceleration (including gravity) at 360 Hz, m/s^2
const LeftTireSetsAvailable = "LeftTireSetsAvailable"                     // How many left tire sets are remaining  255 is unlimited,
const LeftTireSetsUsed = "LeftTireSetsUsed"                               // How many left tire sets used so far,
const LFbrakeLinePress = "LFbrakeLinePress"                               // LF brake line pressure, bar
const LFcoldPressure = "LFcoldPressure"                                   // LF tire cold pressure  as set in the garage, kPa
const LFshockDefl = "LFshockDefl"                                         // LF shock deflection, m
const LFshockDefl_ST = "LFshockDefl_ST"                                   // LF shock deflection at 360 Hz, m
const LFshockVel = "LFshockVel"                                           // LF shock velocity, m/s
const LFshockVel_ST = "LFshockVel_ST"                                     // LF shock velocity at 360 Hz, m/s
const LFtempCL = "LFtempCL"                                               // LF tire left carcass temperature, C
const LFtempCM = "LFtempCM"                                               // LF tire middle carcass temperature, C
const LFtempCR = "LFtempCR"                                               // LF tire right carcass temperature, C
const LFTiresAvailable = "LFTiresAvailable"                               // How many left front tires are remaining  255 is unlimited,
const LFTiresUsed = "LFTiresUsed"                                         // How many left front tires used so far,
const LFwearL = "LFwearL"                                                 // LF tire left percent tread remaining, %
const LFwearM = "LFwearM"                                                 // LF tire middle percent tread remaining, %
const LFwearR = "LFwearR"                                                 // LF tire right percent tread remaining, %
const LoadNumTextures = "LoadNumTextures"                                 // True if the car_num texture will be loaded,
const LongAccel = "LongAccel"                                             // Longitudinal acceleration (including gravity), m/s^2
const LongAccel_ST = "LongAccel_ST"                                       // Longitudinal acceleration (including gravity) at 360 Hz, m/s^2
const LRbrakeLinePress = "LRbrakeLinePress"                               // LR brake line pressure, bar
const LRcoldPressure = "LRcoldPressure"                                   // LR tire cold pressure  as set in the garage, kPa
const LRshockDefl = "LRshockDefl"                                         // LR shock deflection, m
const LRshockDefl_ST = "LRshockDefl_ST"                                   // LR shock deflection at 360 Hz, m
const LRshockVel = "LRshockVel"                                           // LR shock velocity, m/s
const LRshockVel_ST = "LRshockVel_ST"                                     // LR shock velocity at 360 Hz, m/s
const LRtempCL = "LRtempCL"                                               // LR tire left carcass temperature, C
const LRtempCM = "LRtempCM"                                               // LR tire middle carcass temperature, C
const LRtempCR = "LRtempCR"                                               // LR tire right carcass temperature, C
const LRTiresAvailable = "LRTiresAvailable"                               // How many left rear tires are remaining  255 is unlimited,
const LRTiresUsed = "LRTiresUsed"                                         // How many left rear tires used so far,
const LRwearL = "LRwearL"                                                 // LR tire left percent tread remaining, %
const LRwearM = "LRwearM"                                                 // LR tire middle percent tread remaining, %
const LRwearR = "LRwearR"                                                 // LR tire right percent tread remaining, %
const ManifoldPress = "ManifoldPress"                                     // Engine manifold pressure, bar
const ManualBoost = "ManualBoost"                                         // Hybrid manual boost state,
const ManualNoBoost = "ManualNoBoost"                                     // Hybrid manual no boost state,
const MemPageFaultSec = "MemPageFaultSec"                                 // Memory page faults per second,
const MemSoftPageFaultSec = "MemSoftPageFaultSec"                         // Memory soft page faults per second,
const OilLevel = "OilLevel"                                               // Engine oil level, l
const OilPress = "OilPress"                                               // Engine oil pressure, bar
const OilTemp = "OilTemp"                                                 // Engine oil temperature, C
const OkToReloadTextures = "OkToReloadTextures"                           // True if it is ok to reload car textures at this time,
const OnPitRoad = "OnPitRoad"                                             // Is the player car on pit road between the cones,
const PaceMode = "PaceMode"                                               // Are we pacing or not, irsdk_PaceMode
const Pitch = "Pitch"                                                     // Pitch orientation, rad
const PitchRate = "PitchRate"                                             // Pitch rate, rad/s
const PitchRate_ST = "PitchRate_ST"                                       // Pitch rate at 360 Hz, rad/s
const PitOptRepairLeft = "PitOptRepairLeft"                               // Time left for optional repairs if repairs are active, s
const PitRepairLeft = "PitRepairLeft"                                     // Time left for mandatory pit repairs if repairs are active, s
const PitsOpen = "PitsOpen"                                               // True if pit stop is allowed for the current player,
const PitstopActive = "PitstopActive"                                     // Is the player getting pit stop service,
const PitSvFlags = "PitSvFlags"                                           // Bitfield of pit service checkboxes, irsdk_PitSvFlags
const PitSvFuel = "PitSvFuel"                                             // Pit service fuel add amount, l or kWh
const PitSvLFP = "PitSvLFP"                                               // Pit service left front tire pressure, kPa
const PitSvLRP = "PitSvLRP"                                               // Pit service left rear tire pressure, kPa
const PitSvRFP = "PitSvRFP"                                               // Pit service right front tire pressure, kPa
const PitSvRRP = "PitSvRRP"                                               // Pit service right rear tire pressure, kPa
const PitSvTireCompound = "PitSvTireCompound"                             // Pit service pending tire compound,
const PlayerCarClass = "PlayerCarClass"                                   // Player car class id,
const PlayerCarClassPosition = "PlayerCarClassPosition"                   // Players class position in race,
const PlayerCarDriverIncidentCount = "PlayerCarDriverIncidentCount"       // Teams current drivers incident count for this session,
const PlayerCarDryTireSetLimit = "PlayerCarDryTireSetLimit"               // Players dry tire set limit,
const PlayerCarIdx = "PlayerCarIdx"                                       // Players carIdx,
const PlayerCarInPitStall = "PlayerCarInPitStall"                         // Players car is properly in their pitstall,
const PlayerCarMyIncidentCount = "PlayerCarMyIncidentCount"               // Players own incident count for this session,
const PlayerCarPitSvStatus = "PlayerCarPitSvStatus"                       // Players car pit service status bits, irsdk_PitSvStatus
const PlayerCarPosition = "PlayerCarPosition"                             // Players position in race,
const PlayerCarPowerAdjust = "PlayerCarPowerAdjust"                       // Players power adjust, %
const PlayerCarSLBlinkRPM = "PlayerCarSLBlinkRPM"                         // Shift light blink rpm, revs/min
const PlayerCarSLFirstRPM = "PlayerCarSLFirstRPM"                         // Shift light first light rpm, revs/min
const PlayerCarSLLastRPM = "PlayerCarSLLastRPM"                           // Shift light last light rpm, revs/min
const PlayerCarSLShiftRPM = "PlayerCarSLShiftRPM"                         // Shift light shift rpm, revs/min
const PlayerCarTeamIncidentCount = "PlayerCarTeamIncidentCount"           // Players team incident count for this session,
const PlayerCarTowTime = "PlayerCarTowTime"                               // Players car is being towed if time is greater than zero, s
const PlayerCarWeightPenalty = "PlayerCarWeightPenalty"                   // Players weight penalty, kg
const PlayerFastRepairsUsed = "PlayerFastRepairsUsed"                     // Players car number of fast repairs used,
const PlayerTireCompound = "PlayerTireCompound"                           // Players car current tire compound,
const PlayerTrackSurface = "PlayerTrackSurface"                           // Players car track surface type, irsdk_TrkLoc
const PlayerTrackSurfaceMaterial = "PlayerTrackSurfaceMaterial"           // Players car track surface material type, irsdk_TrkSurf
const Precipitation = "Precipitation"                                     // Precipitation at start/finish line, %
const PushToPass = "PushToPass"                                           // Push to pass button state,
const PushToTalk = "PushToTalk"                                           // Push to talk button state,
const RaceLaps = "RaceLaps"                                               // Laps completed in race,
const RadioTransmitCarIdx = "RadioTransmitCarIdx"                         // The car index of the current person speaking on the radio,
const RadioTransmitFrequencyIdx = "RadioTransmitFrequencyIdx"             // The frequency index of the current person speaking on the radio,
const RadioTransmitRadioIdx = "RadioTransmitRadioIdx"                     // The radio index of the current person speaking on the radio,
const RearTireSetsAvailable = "RearTireSetsAvailable"                     // How many rear tire sets are remaining  255 is unlimited,
const RearTireSetsUsed = "RearTireSetsUsed"                               // How many rear tire sets used so far,
const RelativeHumidity = "RelativeHumidity"                               // Relative Humidity at start/finish line, %
const ReplayFrameNum = "ReplayFrameNum"                                   // Integer replay frame number (60 per second),
const ReplayFrameNumEnd = "ReplayFrameNumEnd"                             // Integer replay frame number from end of tape,
const ReplayPlaySlowMotion = "ReplayPlaySlowMotion"                       // 0=not slow motion  1=replay is in slow motion,
const ReplayPlaySpeed = "ReplayPlaySpeed"                                 // Replay playback speed,
const ReplaySessionNum = "ReplaySessionNum"                               // Replay session number,
const ReplaySessionTime = "ReplaySessionTime"                             // Seconds since replay session start, s
const RFbrakeLinePress = "RFbrakeLinePress"                               // RF brake line pressure, bar
const RFcoldPressure = "RFcoldPressure"                                   // RF tire cold pressure  as set in the garage, kPa
const RFshockDefl = "RFshockDefl"                                         // RF shock deflection, m
const RFshockDefl_ST = "RFshockDefl_ST"                                   // RF shock deflection at 360 Hz, m
const RFshockVel = "RFshockVel"                                           // RF shock velocity, m/s
const RFshockVel_ST = "RFshockVel_ST"                                     // RF shock velocity at 360 Hz, m/s
const RFtempCL = "RFtempCL"                                               // RF tire left carcass temperature, C
const RFtempCM = "RFtempCM"                                               // RF tire middle carcass temperature, C
const RFtempCR = "RFtempCR"                                               // RF tire right carcass temperature, C
const RFTiresAvailable = "RFTiresAvailable"                               // How many right front tires are remaining  255 is unlimited,
const RFTiresUsed = "RFTiresUsed"                                         // How many right front tires used so far,
const RFwearL = "RFwearL"                                                 // RF tire left percent tread remaining, %
const RFwearM = "RFwearM"                                                 // RF tire middle percent tread remaining, %
const RFwearR = "RFwearR"                                                 // RF tire right percent tread remaining, %
const RightTireSetsAvailable = "RightTireSetsAvailable"                   // How many right tire sets are remaining  255 is unlimited,
const RightTireSetsUsed = "RightTireSetsUsed"                             // How many right tire sets used so far,
const Roll = "Roll"                                                       // Roll orientation, rad
const RollRate = "RollRate"                                               // Roll rate, rad/s
const RollRate_ST = "RollRate_ST"                                         // Roll rate at 360 Hz, rad/s
const RPM = "RPM"                                                         // Engine rpm, revs/min
const RRbrakeLinePress = "RRbrakeLinePress"                               // RR brake line pressure, bar
const RRcoldPressure = "RRcoldPressure"                                   // RR tire cold pressure  as set in the garage, kPa
const RRshockDefl = "RRshockDefl"                                         // RR shock deflection, m
const RRshockDefl_ST = "RRshockDefl_ST"                                   // RR shock deflection at 360 Hz, m
const RRshockVel = "RRshockVel"                                           // RR shock velocity, m/s
const RRshockVel_ST = "RRshockVel_ST"                                     // RR shock velocity at 360 Hz, m/s
const RRtempCL = "RRtempCL"                                               // RR tire left carcass temperature, C
const RRtempCM = "RRtempCM"                                               // RR tire middle carcass temperature, C
const RRtempCR = "RRtempCR"                                               // RR tire right carcass temperature, C
const RRTiresAvailable = "RRTiresAvailable"                               // How many right rear tires are remaining  255 is unlimited,
const RRTiresUsed = "RRTiresUsed"                                         // How many right rear tires used so far,
const RRwearL = "RRwearL"                                                 // RR tire left percent tread remaining, %
const RRwearM = "RRwearM"                                                 // RR tire middle percent tread remaining, %
const RRwearR = "RRwearR"                                                 // RR tire right percent tread remaining, %
const SessionFlags = "SessionFlags"                                       // Session flags, irsdk_Flags
const SessionJokerLapsRemain = "SessionJokerLapsRemain"                   // Joker laps remaining to be taken,
const SessionLapsRemain = "SessionLapsRemain"                             // Old laps left till session ends use SessionLapsRemainEx,
const SessionLapsRemainEx = "SessionLapsRemainEx"                         // New improved laps left till session ends,
const SessionLapsTotal = "SessionLapsTotal"                               // Total number of laps in session,
const SessionNum = "SessionNum"                                           // Session number,
const SessionOnJokerLap = "SessionOnJokerLap"                             // Player is currently completing a joker lap,
const SessionState = "SessionState"                                       // Session state, irsdk_SessionState
const SessionTick = "SessionTick"                                         // Current update number,
const SessionTime = "SessionTime"                                         // Seconds since session start, s
const SessionTimeOfDay = "SessionTimeOfDay"                               // Time of day in seconds, s
const SessionTimeRemain = "SessionTimeRemain"                             // Seconds left till session ends, s
const SessionTimeTotal = "SessionTimeTotal"                               // Total number of seconds in session, s
const SessionUniqueID = "SessionUniqueID"                                 // Session ID,
const ShiftGrindRPM = "ShiftGrindRPM"                                     // RPM of shifter grinding noise, RPM
const ShiftIndicatorPct = "ShiftIndicatorPct"                             // DEPRECATED use DriverCarSLBlinkRPM instead, %
const ShiftPowerPct = "ShiftPowerPct"                                     // Friction torque applied to gears when shifting or grinding, %
const Skies = "Skies"                                                     // Skies (0=clear/1=p cloudy/2=m cloudy/3=overcast),
const SolarAltitude = "SolarAltitude"                                     // Sun angle above horizon in radians, rad
const SolarAzimuth = "SolarAzimuth"                                       // Sun angle clockwise from north in radians, rad
const Speed = "Speed"                                                     // GPS vehicle speed, m/s
const SteeringWheelAngle = "SteeringWheelAngle"                           // Steering wheel angle, rad
const SteeringWheelAngleMax = "SteeringWheelAngleMax"                     // Steering wheel max angle, rad
const SteeringWheelLimiter = "SteeringWheelLimiter"                       // Force feedback limiter strength limits impacts and oscillation, %
const SteeringWheelMaxForceNm = "SteeringWheelMaxForceNm"                 // Value of strength or max force slider in Nm for FFB, N*m
const SteeringWheelPctDamper = "SteeringWheelPctDamper"                   // Force feedback % max damping, %
const SteeringWheelPctIntensity = "SteeringWheelPctIntensity"             // Force feedback % max intensity, %
const SteeringWheelPctSmoothing = "SteeringWheelPctSmoothing"             // Force feedback % max smoothing, %
const SteeringWheelPctTorque = "SteeringWheelPctTorque"                   // Force feedback % max torque on steering shaft unsigned, %
const SteeringWheelPctTorqueSign = "SteeringWheelPctTorqueSign"           // Force feedback % max torque on steering shaft signed, %
const SteeringWheelPctTorqueSignStops = "SteeringWheelPctTorqueSignStops" // Force feedback % max torque on steering shaft signed stops, %
const SteeringWheelPeakForceNm = "SteeringWheelPeakForceNm"               // Peak torque mapping to direct input units for FFB, N*m
const SteeringWheelTorque = "SteeringWheelTorque"                         // Output torque on steering shaft, N*m
const SteeringWheelTorque_ST = "SteeringWheelTorque_ST"                   // Output torque on steering shaft at 360 Hz, N*m
const SteeringWheelUseLinear = "SteeringWheelUseLinear"                   // True if steering wheel force is using linear mode,
const Throttle = "Throttle"                                               // 0=off throttle to 1=full throttle, %
const ThrottleRaw = "ThrottleRaw"                                         // Raw throttle input 0=off throttle to 1=full throttle, %
const TireLF_RumblePitch = "TireLF_RumblePitch"                           // Players LF Tire Sound rumblestrip pitch, Hz
const TireLR_RumblePitch = "TireLR_RumblePitch"                           // Players LR Tire Sound rumblestrip pitch, Hz
const TireRF_RumblePitch = "TireRF_RumblePitch"                           // Players RF Tire Sound rumblestrip pitch, Hz
const TireRR_RumblePitch = "TireRR_RumblePitch"                           // Players RR Tire Sound rumblestrip pitch, Hz
const TireSetsAvailable = "TireSetsAvailable"                             // How many tire sets are remaining  255 is unlimited,
const TireSetsUsed = "TireSetsUsed"                                       // How many tire sets used so far,
const TrackTemp = "TrackTemp"                                             // Deprecated  set to TrackTempCrew, C
const TrackTempCrew = "TrackTempCrew"                                     // Temperature of track measured by crew around track, C
const VelocityX = "VelocityX"                                             // X velocity, m/s
const VelocityX_ST = "VelocityX_ST"                                       // X velocity, m/s at 360 Hz
const VelocityY = "VelocityY"                                             // Y velocity, m/s
const VelocityY_ST = "VelocityY_ST"                                       // Y velocity, m/s at 360 Hz
const VelocityZ = "VelocityZ"                                             // Z velocity, m/s
const VelocityZ_ST = "VelocityZ_ST"                                       // Z velocity, m/s at 360 Hz
const VertAccel = "VertAccel"                                             // Vertical acceleration (including gravity), m/s^2
const VertAccel_ST = "VertAccel_ST"                                       // Vertical acceleration (including gravity) at 360 Hz, m/s^2
const VidCapActive = "VidCapActive"                                       // True if video currently being captured,
const VidCapEnabled = "VidCapEnabled"                                     // True if video capture system is enabled,
const Voltage = "Voltage"                                                 // Engine voltage, V
const WaterLevel = "WaterLevel"                                           // Engine coolant level, l
const WaterTemp = "WaterTemp"                                             // Engine coolant temp, C
const WeatherType = "WeatherType"                                         // Weather dynamics type, irsdk_WeatherDynamics
const WeatherVersion = "WeatherVersion"                                   // Weather version, irsdk_WeatherVersion
const WindDir = "WindDir"                                                 // Wind direction at start/finish line, rad
const WindVel = "WindVel"                                                 // Wind velocity at start/finish line, m/s
const Yaw = "Yaw"                                                         // Yaw orientation, rad
const YawNorth = "YawNorth"                                               // Yaw orientation relative to north, rad
const YawRate = "YawRate"                                                 // Yaw rate, rad/s
const YawRate_ST = "YawRate_ST"                                           // Yaw rate at 360 Hz, rad/s
