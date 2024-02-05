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
	vars        map[VariableName]Variable
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
	vars := TelemetryVars{vars: make(map[VariableName]Variable, h.numVars)}
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
		vars.vars[VariableName(v.Name)] = v
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

type VariableName string

const AirDensity VariableName = "AirDensity"                                           // Density of air at start/finish line, kg/m^3
const AirPressure VariableName = "AirPressure"                                         // Pressure of air at start/finish line, Pa
const AirTemp VariableName = "AirTemp"                                                 // Temperature of air at start/finish line, C
const Brake VariableName = "Brake"                                                     // 0=brake released to 1=max pedal force, %
const BrakeABSactive VariableName = "BrakeABSactive"                                   // true if abs is currently reducing brake force pressure,
const BrakeRaw VariableName = "BrakeRaw"                                               // Raw brake input 0=brake released to 1=max pedal force, %
const CamCameraNumber VariableName = "CamCameraNumber"                                 // Active camera number,
const CamCameraState VariableName = "CamCameraState"                                   // State of camera system, irsdk_CameraState
const CamCarIdx VariableName = "CamCarIdx"                                             // Active camera's focus car index,
const CamGroupNumber VariableName = "CamGroupNumber"                                   // Active camera group number,
const CarIdxBestLapNum VariableName = "CarIdxBestLapNum"                               // Cars best lap number,
const CarIdxBestLapTime VariableName = "CarIdxBestLapTime"                             // Cars best lap time, s
const CarIdxClass VariableName = "CarIdxClass"                                         // Cars class id by car index,
const CarIdxClassPosition VariableName = "CarIdxClassPosition"                         // Cars class position in race by car index,
const CarIdxEstTime VariableName = "CarIdxEstTime"                                     // Estimated time to reach current location on track, s
const CarIdxF2Time VariableName = "CarIdxF2Time"                                       // Race time behind leader or fastest lap time otherwise, s
const CarIdxFastRepairsUsed VariableName = "CarIdxFastRepairsUsed"                     // How many fast repairs each car has used,
const CarIdxGear VariableName = "CarIdxGear"                                           // -1=reverse  0=neutral  1..n=current gear by car index,
const CarIdxLap VariableName = "CarIdxLap"                                             // Laps started by car index,
const CarIdxLapCompleted VariableName = "CarIdxLapCompleted"                           // Laps completed by car index,
const CarIdxLapDistPct VariableName = "CarIdxLapDistPct"                               // Percentage distance around lap by car index, %
const CarIdxLastLapTime VariableName = "CarIdxLastLapTime"                             // Cars last lap time, s
const CarIdxOnPitRoad VariableName = "CarIdxOnPitRoad"                                 // On pit road between the cones by car index,
const CarIdxP2P_Count VariableName = "CarIdxP2P_Count"                                 // Push2Pass count of usage (or remaining in Race),
const CarIdxP2P_Status VariableName = "CarIdxP2P_Status"                               // Push2Pass active or not,
const CarIdxPaceFlags VariableName = "CarIdxPaceFlags"                                 // Pacing status flags for each car, irsdk_PaceFlags
const CarIdxPaceLine VariableName = "CarIdxPaceLine"                                   // What line cars are pacing in  or -1 if not pacing,
const CarIdxPaceRow VariableName = "CarIdxPaceRow"                                     // What row cars are pacing in  or -1 if not pacing,
const CarIdxPosition VariableName = "CarIdxPosition"                                   // Cars position in race by car index,
const CarIdxQualTireCompound VariableName = "CarIdxQualTireCompound"                   // Cars Qual tire compound,
const CarIdxQualTireCompoundLocked VariableName = "CarIdxQualTireCompoundLocked"       // Cars Qual tire compound is locked-in,
const CarIdxRPM VariableName = "CarIdxRPM"                                             // Engine rpm by car index, revs/min
const CarIdxSessionFlags VariableName = "CarIdxSessionFlags"                           // Session flags for each player, irsdk_Flags
const CarIdxSteer VariableName = "CarIdxSteer"                                         // Steering wheel angle by car index, rad
const CarIdxTireCompound VariableName = "CarIdxTireCompound"                           // Cars current tire compound,
const CarIdxTrackSurface VariableName = "CarIdxTrackSurface"                           // Track surface type by car index, irsdk_TrkLoc
const CarIdxTrackSurfaceMaterial VariableName = "CarIdxTrackSurfaceMaterial"           // Track surface material type by car index, irsdk_TrkSurf
const CarLeftRight VariableName = "CarLeftRight"                                       // Notify if car is to the left or right of driver, irsdk_CarLeftRight
const ChanAvgLatency VariableName = "ChanAvgLatency"                                   // Communications average latency, s
const ChanClockSkew VariableName = "ChanClockSkew"                                     // Communications server clock skew, s
const ChanLatency VariableName = "ChanLatency"                                         // Communications latency, s
const ChanPartnerQuality VariableName = "ChanPartnerQuality"                           // Partner communications quality, %
const ChanQuality VariableName = "ChanQuality"                                         // Communications quality, %
const Clutch VariableName = "Clutch"                                                   // 0=disengaged to 1=fully engaged, %
const ClutchRaw VariableName = "ClutchRaw"                                             // Raw clutch input 0=disengaged to 1=fully engaged, %
const CpuUsageBG VariableName = "CpuUsageBG"                                           // Percent of available tim bg thread took with a 1 sec avg, %
const CpuUsageFG VariableName = "CpuUsageFG"                                           // Percent of available tim fg thread took with a 1 sec avg, %
const DCDriversSoFar VariableName = "DCDriversSoFar"                                   // Number of team drivers who have run a stint,
const DCLapStatus VariableName = "DCLapStatus"                                         // Status of driver change lap requirements,
const dcPitSpeedLimiterToggle VariableName = "dcPitSpeedLimiterToggle"                 // In car traction control active,
const dcStarter VariableName = "dcStarter"                                             // In car trigger car starter,
const DisplayUnits VariableName = "DisplayUnits"                                       // Default units for the user interface 0 = english 1  VariableName = metric,
const dpFastRepair VariableName = "dpFastRepair"                                       // Pitstop fast repair set,
const dpFuelAddKg VariableName = "dpFuelAddKg"                                         // Pitstop fuel add amount, kg
const dpFuelAutoFillActive VariableName = "dpFuelAutoFillActive"                       // Pitstop auto fill fuel next stop flag,
const dpFuelAutoFillEnabled VariableName = "dpFuelAutoFillEnabled"                     // Pitstop auto fill fuel system enabled,
const dpFuelFill VariableName = "dpFuelFill"                                           // Pitstop fuel fill flag,
const dpLFTireChange VariableName = "dpLFTireChange"                                   // Pitstop lf tire change request,
const dpLFTireColdPress VariableName = "dpLFTireColdPress"                             // Pitstop lf tire cold pressure adjustment, Pa
const dpLRTireChange VariableName = "dpLRTireChange"                                   // Pitstop lr tire change request,
const dpLRTireColdPress VariableName = "dpLRTireColdPress"                             // Pitstop lr tire cold pressure adjustment, Pa
const dpRFTireChange VariableName = "dpRFTireChange"                                   // Pitstop rf tire change request,
const dpRFTireColdPress VariableName = "dpRFTireColdPress"                             // Pitstop rf cold tire pressure adjustment, Pa
const dpRRTireChange VariableName = "dpRRTireChange"                                   // Pitstop rr tire change request,
const dpRRTireColdPress VariableName = "dpRRTireColdPress"                             // Pitstop rr cold tire pressure adjustment, Pa
const dpWindshieldTearoff VariableName = "dpWindshieldTearoff"                         // Pitstop windshield tearoff,
const DriverMarker VariableName = "DriverMarker"                                       // Driver activated flag,
const Engine0_RPM VariableName = "Engine0_RPM"                                         // Engine0Engine rpm, revs/min
const EngineWarnings VariableName = "EngineWarnings"                                   // Bitfield for warning lights, irsdk_EngineWarnings
const EnterExitReset VariableName = "EnterExitReset"                                   // Indicate action the reset key will take 0 enter 1 exit 2 reset,
const FastRepairAvailable VariableName = "FastRepairAvailable"                         // How many fast repairs left  255 is unlimited,
const FastRepairUsed VariableName = "FastRepairUsed"                                   // How many fast repairs used so far,
const FogLevel VariableName = "FogLevel"                                               // Fog level at start/finish line, %
const FrameRate VariableName = "FrameRate"                                             // Average frames per second, fps
const FrontTireSetsAvailable VariableName = "FrontTireSetsAvailable"                   // How many front tire sets are remaining  255 is unlimited,
const FrontTireSetsUsed VariableName = "FrontTireSetsUsed"                             // How many front tire sets used so far,
const FuelLevel VariableName = "FuelLevel"                                             // Liters of fuel remaining, l
const FuelLevelPct VariableName = "FuelLevelPct"                                       // Percent fuel remaining, %
const FuelPress VariableName = "FuelPress"                                             // Engine fuel pressure, bar
const FuelUsePerHour VariableName = "FuelUsePerHour"                                   // Engine fuel used instantaneous, kg/h
const Gear VariableName = "Gear"                                                       // -1=reverse  0=neutral  1..n=current gear,
const GpuUsage VariableName = "GpuUsage"                                               // Percent of available tim gpu took with a 1 sec avg, %
const HandbrakeRaw VariableName = "HandbrakeRaw"                                       // Raw handbrake input 0=handbrake released to 1=max force, %
const IsDiskLoggingActive VariableName = "IsDiskLoggingActive"                         // 0=disk based telemetry file not being written  1=being written,
const IsDiskLoggingEnabled VariableName = "IsDiskLoggingEnabled"                       // 0=disk based telemetry turned off  1=turned on,
const IsGarageVisible VariableName = "IsGarageVisible"                                 // 1=Garage screen is visible,
const IsInGarage VariableName = "IsInGarage"                                           // 1=Car in garage physics running,
const IsOnTrack VariableName = "IsOnTrack"                                             // 1=Car on track physics running with player in car,
const IsOnTrackCar VariableName = "IsOnTrackCar"                                       // 1=Car on track physics running,
const IsReplayPlaying VariableName = "IsReplayPlaying"                                 // 0=replay not playing  1=replay playing,
const Lap VariableName = "Lap"                                                         // Laps started count,
const LapBestLap VariableName = "LapBestLap"                                           // Players best lap number,
const LapBestLapTime VariableName = "LapBestLapTime"                                   // Players best lap time, s
const LapBestNLapLap VariableName = "LapBestNLapLap"                                   // Player last lap in best N average lap time,
const LapBestNLapTime VariableName = "LapBestNLapTime"                                 // Player best N average lap time, s
const LapCompleted VariableName = "LapCompleted"                                       // Laps completed count,
const LapCurrentLapTime VariableName = "LapCurrentLapTime"                             // Estimate of players current lap time as shown in F3 box, s
const LapDeltaToBestLap VariableName = "LapDeltaToBestLap"                             // Delta time for best lap, s
const LapDeltaToBestLap_DD VariableName = "LapDeltaToBestLap_DD"                       // Rate of change of delta time for best lap, s/s
const LapDeltaToBestLap_OK VariableName = "LapDeltaToBestLap_OK"                       // Delta time for best lap is valid,
const LapDeltaToOptimalLap VariableName = "LapDeltaToOptimalLap"                       // Delta time for optimal lap, s
const LapDeltaToOptimalLap_DD VariableName = "LapDeltaToOptimalLap_DD"                 // Rate of change of delta time for optimal lap, s/s
const LapDeltaToOptimalLap_OK VariableName = "LapDeltaToOptimalLap_OK"                 // Delta time for optimal lap is valid,
const LapDeltaToSessionBestLap VariableName = "LapDeltaToSessionBestLap"               // Delta time for session best lap, s
const LapDeltaToSessionBestLap_DD VariableName = "LapDeltaToSessionBestLap_DD"         // Rate of change of delta time for session best lap, s/s
const LapDeltaToSessionBestLap_OK VariableName = "LapDeltaToSessionBestLap_OK"         // Delta time for session best lap is valid,
const LapDeltaToSessionLastlLap VariableName = "LapDeltaToSessionLastlLap"             // Delta time for session last lap, s
const LapDeltaToSessionLastlLap_DD VariableName = "LapDeltaToSessionLastlLap_DD"       // Rate of change of delta time for session last lap, s/s
const LapDeltaToSessionLastlLap_OK VariableName = "LapDeltaToSessionLastlLap_OK"       // Delta time for session last lap is valid,
const LapDeltaToSessionOptimalLap VariableName = "LapDeltaToSessionOptimalLap"         // Delta time for session optimal lap, s
const LapDeltaToSessionOptimalLap_DD VariableName = "LapDeltaToSessionOptimalLap_DD"   // Rate of change of delta time for session optimal lap, s/s
const LapDeltaToSessionOptimalLap_OK VariableName = "LapDeltaToSessionOptimalLap_OK"   // Delta time for session optimal lap is valid,
const LapDist VariableName = "LapDist"                                                 // Meters traveled from S/F this lap, m
const LapDistPct VariableName = "LapDistPct"                                           // Percentage distance around lap, %
const LapLasNLapSeq VariableName = "LapLasNLapSeq"                                     // Player num consecutive clean laps completed for N average,
const LapLastLapTime VariableName = "LapLastLapTime"                                   // Players last lap time, s
const LapLastNLapTime VariableName = "LapLastNLapTime"                                 // Player last N average lap time, s
const LatAccel VariableName = "LatAccel"                                               // Lateral acceleration (including gravity), m/s^2
const LatAccel_ST VariableName = "LatAccel_ST"                                         // Lateral acceleration (including gravity) at 360 Hz, m/s^2
const LeftTireSetsAvailable VariableName = "LeftTireSetsAvailable"                     // How many left tire sets are remaining  255 is unlimited,
const LeftTireSetsUsed VariableName = "LeftTireSetsUsed"                               // How many left tire sets used so far,
const LFbrakeLinePress VariableName = "LFbrakeLinePress"                               // LF brake line pressure, bar
const LFcoldPressure VariableName = "LFcoldPressure"                                   // LF tire cold pressure  as set in the garage, kPa
const LFshockDefl VariableName = "LFshockDefl"                                         // LF shock deflection, m
const LFshockDefl_ST VariableName = "LFshockDefl_ST"                                   // LF shock deflection at 360 Hz, m
const LFshockVel VariableName = "LFshockVel"                                           // LF shock velocity, m/s
const LFshockVel_ST VariableName = "LFshockVel_ST"                                     // LF shock velocity at 360 Hz, m/s
const LFtempCL VariableName = "LFtempCL"                                               // LF tire left carcass temperature, C
const LFtempCM VariableName = "LFtempCM"                                               // LF tire middle carcass temperature, C
const LFtempCR VariableName = "LFtempCR"                                               // LF tire right carcass temperature, C
const LFTiresAvailable VariableName = "LFTiresAvailable"                               // How many left front tires are remaining  255 is unlimited,
const LFTiresUsed VariableName = "LFTiresUsed"                                         // How many left front tires used so far,
const LFwearL VariableName = "LFwearL"                                                 // LF tire left percent tread remaining, %
const LFwearM VariableName = "LFwearM"                                                 // LF tire middle percent tread remaining, %
const LFwearR VariableName = "LFwearR"                                                 // LF tire right percent tread remaining, %
const LoadNumTextures VariableName = "LoadNumTextures"                                 // True if the car_num texture will be loaded,
const LongAccel VariableName = "LongAccel"                                             // Longitudinal acceleration (including gravity), m/s^2
const LongAccel_ST VariableName = "LongAccel_ST"                                       // Longitudinal acceleration (including gravity) at 360 Hz, m/s^2
const LRbrakeLinePress VariableName = "LRbrakeLinePress"                               // LR brake line pressure, bar
const LRcoldPressure VariableName = "LRcoldPressure"                                   // LR tire cold pressure  as set in the garage, kPa
const LRshockDefl VariableName = "LRshockDefl"                                         // LR shock deflection, m
const LRshockDefl_ST VariableName = "LRshockDefl_ST"                                   // LR shock deflection at 360 Hz, m
const LRshockVel VariableName = "LRshockVel"                                           // LR shock velocity, m/s
const LRshockVel_ST VariableName = "LRshockVel_ST"                                     // LR shock velocity at 360 Hz, m/s
const LRtempCL VariableName = "LRtempCL"                                               // LR tire left carcass temperature, C
const LRtempCM VariableName = "LRtempCM"                                               // LR tire middle carcass temperature, C
const LRtempCR VariableName = "LRtempCR"                                               // LR tire right carcass temperature, C
const LRTiresAvailable VariableName = "LRTiresAvailable"                               // How many left rear tires are remaining  255 is unlimited,
const LRTiresUsed VariableName = "LRTiresUsed"                                         // How many left rear tires used so far,
const LRwearL VariableName = "LRwearL"                                                 // LR tire left percent tread remaining, %
const LRwearM VariableName = "LRwearM"                                                 // LR tire middle percent tread remaining, %
const LRwearR VariableName = "LRwearR"                                                 // LR tire right percent tread remaining, %
const ManifoldPress VariableName = "ManifoldPress"                                     // Engine manifold pressure, bar
const ManualBoost VariableName = "ManualBoost"                                         // Hybrid manual boost state,
const ManualNoBoost VariableName = "ManualNoBoost"                                     // Hybrid manual no boost state,
const MemPageFaultSec VariableName = "MemPageFaultSec"                                 // Memory page faults per second,
const MemSoftPageFaultSec VariableName = "MemSoftPageFaultSec"                         // Memory soft page faults per second,
const OilLevel VariableName = "OilLevel"                                               // Engine oil level, l
const OilPress VariableName = "OilPress"                                               // Engine oil pressure, bar
const OilTemp VariableName = "OilTemp"                                                 // Engine oil temperature, C
const OkToReloadTextures VariableName = "OkToReloadTextures"                           // True if it is ok to reload car textures at this time,
const OnPitRoad VariableName = "OnPitRoad"                                             // Is the player car on pit road between the cones,
const PaceMode VariableName = "PaceMode"                                               // Are we pacing or not, irsdk_PaceMode
const Pitch VariableName = "Pitch"                                                     // Pitch orientation, rad
const PitchRate VariableName = "PitchRate"                                             // Pitch rate, rad/s
const PitchRate_ST VariableName = "PitchRate_ST"                                       // Pitch rate at 360 Hz, rad/s
const PitOptRepairLeft VariableName = "PitOptRepairLeft"                               // Time left for optional repairs if repairs are active, s
const PitRepairLeft VariableName = "PitRepairLeft"                                     // Time left for mandatory pit repairs if repairs are active, s
const PitsOpen VariableName = "PitsOpen"                                               // True if pit stop is allowed for the current player,
const PitstopActive VariableName = "PitstopActive"                                     // Is the player getting pit stop service,
const PitSvFlags VariableName = "PitSvFlags"                                           // Bitfield of pit service checkboxes, irsdk_PitSvFlags
const PitSvFuel VariableName = "PitSvFuel"                                             // Pit service fuel add amount, l or kWh
const PitSvLFP VariableName = "PitSvLFP"                                               // Pit service left front tire pressure, kPa
const PitSvLRP VariableName = "PitSvLRP"                                               // Pit service left rear tire pressure, kPa
const PitSvRFP VariableName = "PitSvRFP"                                               // Pit service right front tire pressure, kPa
const PitSvRRP VariableName = "PitSvRRP"                                               // Pit service right rear tire pressure, kPa
const PitSvTireCompound VariableName = "PitSvTireCompound"                             // Pit service pending tire compound,
const PlayerCarClass VariableName = "PlayerCarClass"                                   // Player car class id,
const PlayerCarClassPosition VariableName = "PlayerCarClassPosition"                   // Players class position in race,
const PlayerCarDriverIncidentCount VariableName = "PlayerCarDriverIncidentCount"       // Teams current drivers incident count for this session,
const PlayerCarDryTireSetLimit VariableName = "PlayerCarDryTireSetLimit"               // Players dry tire set limit,
const PlayerCarIdx VariableName = "PlayerCarIdx"                                       // Players carIdx,
const PlayerCarInPitStall VariableName = "PlayerCarInPitStall"                         // Players car is properly in their pitstall,
const PlayerCarMyIncidentCount VariableName = "PlayerCarMyIncidentCount"               // Players own incident count for this session,
const PlayerCarPitSvStatus VariableName = "PlayerCarPitSvStatus"                       // Players car pit service status bits, irsdk_PitSvStatus
const PlayerCarPosition VariableName = "PlayerCarPosition"                             // Players position in race,
const PlayerCarPowerAdjust VariableName = "PlayerCarPowerAdjust"                       // Players power adjust, %
const PlayerCarSLBlinkRPM VariableName = "PlayerCarSLBlinkRPM"                         // Shift light blink rpm, revs/min
const PlayerCarSLFirstRPM VariableName = "PlayerCarSLFirstRPM"                         // Shift light first light rpm, revs/min
const PlayerCarSLLastRPM VariableName = "PlayerCarSLLastRPM"                           // Shift light last light rpm, revs/min
const PlayerCarSLShiftRPM VariableName = "PlayerCarSLShiftRPM"                         // Shift light shift rpm, revs/min
const PlayerCarTeamIncidentCount VariableName = "PlayerCarTeamIncidentCount"           // Players team incident count for this session,
const PlayerCarTowTime VariableName = "PlayerCarTowTime"                               // Players car is being towed if time is greater than zero, s
const PlayerCarWeightPenalty VariableName = "PlayerCarWeightPenalty"                   // Players weight penalty, kg
const PlayerFastRepairsUsed VariableName = "PlayerFastRepairsUsed"                     // Players car number of fast repairs used,
const PlayerTireCompound VariableName = "PlayerTireCompound"                           // Players car current tire compound,
const PlayerTrackSurface VariableName = "PlayerTrackSurface"                           // Players car track surface type, irsdk_TrkLoc
const PlayerTrackSurfaceMaterial VariableName = "PlayerTrackSurfaceMaterial"           // Players car track surface material type, irsdk_TrkSurf
const Precipitation VariableName = "Precipitation"                                     // Precipitation at start/finish line, %
const PushToPass VariableName = "PushToPass"                                           // Push to pass button state,
const PushToTalk VariableName = "PushToTalk"                                           // Push to talk button state,
const RaceLaps VariableName = "RaceLaps"                                               // Laps completed in race,
const RadioTransmitCarIdx VariableName = "RadioTransmitCarIdx"                         // The car index of the current person speaking on the radio,
const RadioTransmitFrequencyIdx VariableName = "RadioTransmitFrequencyIdx"             // The frequency index of the current person speaking on the radio,
const RadioTransmitRadioIdx VariableName = "RadioTransmitRadioIdx"                     // The radio index of the current person speaking on the radio,
const RearTireSetsAvailable VariableName = "RearTireSetsAvailable"                     // How many rear tire sets are remaining  255 is unlimited,
const RearTireSetsUsed VariableName = "RearTireSetsUsed"                               // How many rear tire sets used so far,
const RelativeHumidity VariableName = "RelativeHumidity"                               // Relative Humidity at start/finish line, %
const ReplayFrameNum VariableName = "ReplayFrameNum"                                   // Integer replay frame number (60 per second),
const ReplayFrameNumEnd VariableName = "ReplayFrameNumEnd"                             // Integer replay frame number from end of tape,
const ReplayPlaySlowMotion VariableName = "ReplayPlaySlowMotion"                       // 0=not slow motion  1=replay is in slow motion,
const ReplayPlaySpeed VariableName = "ReplayPlaySpeed"                                 // Replay playback speed,
const ReplaySessionNum VariableName = "ReplaySessionNum"                               // Replay session number,
const ReplaySessionTime VariableName = "ReplaySessionTime"                             // Seconds since replay session start, s
const RFbrakeLinePress VariableName = "RFbrakeLinePress"                               // RF brake line pressure, bar
const RFcoldPressure VariableName = "RFcoldPressure"                                   // RF tire cold pressure  as set in the garage, kPa
const RFshockDefl VariableName = "RFshockDefl"                                         // RF shock deflection, m
const RFshockDefl_ST VariableName = "RFshockDefl_ST"                                   // RF shock deflection at 360 Hz, m
const RFshockVel VariableName = "RFshockVel"                                           // RF shock velocity, m/s
const RFshockVel_ST VariableName = "RFshockVel_ST"                                     // RF shock velocity at 360 Hz, m/s
const RFtempCL VariableName = "RFtempCL"                                               // RF tire left carcass temperature, C
const RFtempCM VariableName = "RFtempCM"                                               // RF tire middle carcass temperature, C
const RFtempCR VariableName = "RFtempCR"                                               // RF tire right carcass temperature, C
const RFTiresAvailable VariableName = "RFTiresAvailable"                               // How many right front tires are remaining  255 is unlimited,
const RFTiresUsed VariableName = "RFTiresUsed"                                         // How many right front tires used so far,
const RFwearL VariableName = "RFwearL"                                                 // RF tire left percent tread remaining, %
const RFwearM VariableName = "RFwearM"                                                 // RF tire middle percent tread remaining, %
const RFwearR VariableName = "RFwearR"                                                 // RF tire right percent tread remaining, %
const RightTireSetsAvailable VariableName = "RightTireSetsAvailable"                   // How many right tire sets are remaining  255 is unlimited,
const RightTireSetsUsed VariableName = "RightTireSetsUsed"                             // How many right tire sets used so far,
const Roll VariableName = "Roll"                                                       // Roll orientation, rad
const RollRate VariableName = "RollRate"                                               // Roll rate, rad/s
const RollRate_ST VariableName = "RollRate_ST"                                         // Roll rate at 360 Hz, rad/s
const RPM VariableName = "RPM"                                                         // Engine rpm, revs/min
const RRbrakeLinePress VariableName = "RRbrakeLinePress"                               // RR brake line pressure, bar
const RRcoldPressure VariableName = "RRcoldPressure"                                   // RR tire cold pressure  as set in the garage, kPa
const RRshockDefl VariableName = "RRshockDefl"                                         // RR shock deflection, m
const RRshockDefl_ST VariableName = "RRshockDefl_ST"                                   // RR shock deflection at 360 Hz, m
const RRshockVel VariableName = "RRshockVel"                                           // RR shock velocity, m/s
const RRshockVel_ST VariableName = "RRshockVel_ST"                                     // RR shock velocity at 360 Hz, m/s
const RRtempCL VariableName = "RRtempCL"                                               // RR tire left carcass temperature, C
const RRtempCM VariableName = "RRtempCM"                                               // RR tire middle carcass temperature, C
const RRtempCR VariableName = "RRtempCR"                                               // RR tire right carcass temperature, C
const RRTiresAvailable VariableName = "RRTiresAvailable"                               // How many right rear tires are remaining  255 is unlimited,
const RRTiresUsed VariableName = "RRTiresUsed"                                         // How many right rear tires used so far,
const RRwearL VariableName = "RRwearL"                                                 // RR tire left percent tread remaining, %
const RRwearM VariableName = "RRwearM"                                                 // RR tire middle percent tread remaining, %
const RRwearR VariableName = "RRwearR"                                                 // RR tire right percent tread remaining, %
const SessionFlags VariableName = "SessionFlags"                                       // Session flags, irsdk_Flags
const SessionJokerLapsRemain VariableName = "SessionJokerLapsRemain"                   // Joker laps remaining to be taken,
const SessionLapsRemain VariableName = "SessionLapsRemain"                             // Old laps left till session ends use SessionLapsRemainEx,
const SessionLapsRemainEx VariableName = "SessionLapsRemainEx"                         // New improved laps left till session ends,
const SessionLapsTotal VariableName = "SessionLapsTotal"                               // Total number of laps in session,
const SessionNum VariableName = "SessionNum"                                           // Session number,
const SessionOnJokerLap VariableName = "SessionOnJokerLap"                             // Player is currently completing a joker lap,
const SessionState VariableName = "SessionState"                                       // Session state, irsdk_SessionState
const SessionTick VariableName = "SessionTick"                                         // Current update number,
const SessionTime VariableName = "SessionTime"                                         // Seconds since session start, s
const SessionTimeOfDay VariableName = "SessionTimeOfDay"                               // Time of day in seconds, s
const SessionTimeRemain VariableName = "SessionTimeRemain"                             // Seconds left till session ends, s
const SessionTimeTotal VariableName = "SessionTimeTotal"                               // Total number of seconds in session, s
const SessionUniqueID VariableName = "SessionUniqueID"                                 // Session ID,
const ShiftGrindRPM VariableName = "ShiftGrindRPM"                                     // RPM of shifter grinding noise, RPM
const ShiftIndicatorPct VariableName = "ShiftIndicatorPct"                             // DEPRECATED use DriverCarSLBlinkRPM instead, %
const ShiftPowerPct VariableName = "ShiftPowerPct"                                     // Friction torque applied to gears when shifting or grinding, %
const Skies VariableName = "Skies"                                                     // Skies (0=clear/1=p cloudy/2=m cloudy/3=overcast),
const SolarAltitude VariableName = "SolarAltitude"                                     // Sun angle above horizon in radians, rad
const SolarAzimuth VariableName = "SolarAzimuth"                                       // Sun angle clockwise from north in radians, rad
const Speed VariableName = "Speed"                                                     // GPS vehicle speed, m/s
const SteeringWheelAngle VariableName = "SteeringWheelAngle"                           // Steering wheel angle, rad
const SteeringWheelAngleMax VariableName = "SteeringWheelAngleMax"                     // Steering wheel max angle, rad
const SteeringWheelLimiter VariableName = "SteeringWheelLimiter"                       // Force feedback limiter strength limits impacts and oscillation, %
const SteeringWheelMaxForceNm VariableName = "SteeringWheelMaxForceNm"                 // Value of strength or max force slider in Nm for FFB, N*m
const SteeringWheelPctDamper VariableName = "SteeringWheelPctDamper"                   // Force feedback % max damping, %
const SteeringWheelPctIntensity VariableName = "SteeringWheelPctIntensity"             // Force feedback % max intensity, %
const SteeringWheelPctSmoothing VariableName = "SteeringWheelPctSmoothing"             // Force feedback % max smoothing, %
const SteeringWheelPctTorque VariableName = "SteeringWheelPctTorque"                   // Force feedback % max torque on steering shaft unsigned, %
const SteeringWheelPctTorqueSign VariableName = "SteeringWheelPctTorqueSign"           // Force feedback % max torque on steering shaft signed, %
const SteeringWheelPctTorqueSignStops VariableName = "SteeringWheelPctTorqueSignStops" // Force feedback % max torque on steering shaft signed stops, %
const SteeringWheelPeakForceNm VariableName = "SteeringWheelPeakForceNm"               // Peak torque mapping to direct input units for FFB, N*m
const SteeringWheelTorque VariableName = "SteeringWheelTorque"                         // Output torque on steering shaft, N*m
const SteeringWheelTorque_ST VariableName = "SteeringWheelTorque_ST"                   // Output torque on steering shaft at 360 Hz, N*m
const SteeringWheelUseLinear VariableName = "SteeringWheelUseLinear"                   // True if steering wheel force is using linear mode,
const Throttle VariableName = "Throttle"                                               // 0=off throttle to 1=full throttle, %
const ThrottleRaw VariableName = "ThrottleRaw"                                         // Raw throttle input 0=off throttle to 1=full throttle, %
const TireLF_RumblePitch VariableName = "TireLF_RumblePitch"                           // Players LF Tire Sound rumblestrip pitch, Hz
const TireLR_RumblePitch VariableName = "TireLR_RumblePitch"                           // Players LR Tire Sound rumblestrip pitch, Hz
const TireRF_RumblePitch VariableName = "TireRF_RumblePitch"                           // Players RF Tire Sound rumblestrip pitch, Hz
const TireRR_RumblePitch VariableName = "TireRR_RumblePitch"                           // Players RR Tire Sound rumblestrip pitch, Hz
const TireSetsAvailable VariableName = "TireSetsAvailable"                             // How many tire sets are remaining  255 is unlimited,
const TireSetsUsed VariableName = "TireSetsUsed"                                       // How many tire sets used so far,
const TrackTemp VariableName = "TrackTemp"                                             // Deprecated  set to TrackTempCrew, C
const TrackTempCrew VariableName = "TrackTempCrew"                                     // Temperature of track measured by crew around track, C
const VelocityX VariableName = "VelocityX"                                             // X velocity, m/s
const VelocityX_ST VariableName = "VelocityX_ST"                                       // X velocity, m/s at 360 Hz
const VelocityY VariableName = "VelocityY"                                             // Y velocity, m/s
const VelocityY_ST VariableName = "VelocityY_ST"                                       // Y velocity, m/s at 360 Hz
const VelocityZ VariableName = "VelocityZ"                                             // Z velocity, m/s
const VelocityZ_ST VariableName = "VelocityZ_ST"                                       // Z velocity, m/s at 360 Hz
const VertAccel VariableName = "VertAccel"                                             // Vertical acceleration (including gravity), m/s^2
const VertAccel_ST VariableName = "VertAccel_ST"                                       // Vertical acceleration (including gravity) at 360 Hz, m/s^2
const VidCapActive VariableName = "VidCapActive"                                       // True if video currently being captured,
const VidCapEnabled VariableName = "VidCapEnabled"                                     // True if video capture system is enabled,
const Voltage VariableName = "Voltage"                                                 // Engine voltage, V
const WaterLevel VariableName = "WaterLevel"                                           // Engine coolant level, l
const WaterTemp VariableName = "WaterTemp"                                             // Engine coolant temp, C
const WeatherType VariableName = "WeatherType"                                         // Weather dynamics type, irsdk_WeatherDynamics
const WeatherVersion VariableName = "WeatherVersion"                                   // Weather version, irsdk_WeatherVersion
const WindDir VariableName = "WindDir"                                                 // Wind direction at start/finish line, rad
const WindVel VariableName = "WindVel"                                                 // Wind velocity at start/finish line, m/s
const Yaw VariableName = "Yaw"                                                         // Yaw orientation, rad
const YawNorth VariableName = "YawNorth"                                               // Yaw orientation relative to north, rad
const YawRate VariableName = "YawRate"                                                 // Yaw rate, rad/s
const YawRate_ST VariableName = "YawRate_ST"                                           // Yaw rate at 360 Hz, rad/s
