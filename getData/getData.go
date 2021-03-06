package getdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Client HTTPClient

//HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//sets client to instance of httpclient when it initialises. Init function runs once
//when package is imported
func init() {
	Client = &http.Client{}
}

//Get will get the data from espn and return a struct of info we want. What is that info?
func Get() string {

	// want to return EspnData.Centre.Match.CurrentSummary

	client := Client

	req, err := http.NewRequest("GET", `https://www.espncricinfo.com/matches/engine/match/1198241.json`, nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	s := new(EspnData)

	err = json.Unmarshal(body, &s)

	return s.Match.CurrentSummary

}

type EspnData struct {
	Centre struct {
		Batting []struct {
			BallsFaced        string `json:"balls_faced"`
			BattingStyle      string `json:"batting_style"`
			ControlPercentage string `json:"control_percentage"`
			DismissalName     string `json:"dismissal_name"`
			DotBallPercentage string `json:"dot_ball_percentage"`
			KnownAs           string `json:"known_as"`
			LiveCurrentName   string `json:"live_current_name"`
			MatchAward        int    `json:"match_award"`
			Notout            string `json:"notout"`
			PlayerID          string `json:"player_id"`
			PopularName       string `json:"popular_name"`
			PreferredShot     struct {
				BallsFaced  string   `json:"balls_faced"`
				Runs        string   `json:"runs"`
				RunsSummary []string `json:"runs_summary"`
				ShotName    string   `json:"shot_name"`
			} `json:"preferred_shot"`
			Runs         string   `json:"runs"`
			RunsSummary  []string `json:"runs_summary"`
			ScoringShots string   `json:"scoring_shots"`
			StrikeRate   string   `json:"strike_rate"`
			WagonZone    []struct {
				Runs         int   `json:"runs"`
				RunsSummary  []int `json:"runs_summary"`
				ScoringShots int   `json:"scoring_shots"`
			} `json:"wagon_zone"`
		} `json:"batting"`
		Bowling []struct {
			BowlingStyle    string `json:"bowling_style"`
			Conceded        string `json:"conceded"`
			EconomyRate     string `json:"economy_rate"`
			KnownAs         string `json:"known_as"`
			LiveCurrentName string `json:"live_current_name"`
			Maidens         string `json:"maidens"`
			MatchAward      int    `json:"match_award"`
			OverallLhb      struct {
				Balls       string `json:"balls"`
				Conceded    string `json:"conceded"`
				EconomyRate string `json:"economy_rate"`
				Wickets     string `json:"wickets"`
			} `json:"overall_lhb"`
			OverallRhb struct {
				Balls       string `json:"balls"`
				Conceded    string `json:"conceded"`
				EconomyRate string `json:"economy_rate"`
				Wickets     string `json:"wickets"`
			} `json:"overall_rhb"`
			Overs       string    `json:"overs"`
			PitchMapLhb [][][]int `json:"pitch_map_lhb"`
			PitchMapRhb [][][]int `json:"pitch_map_rhb"`
			PlayerID    string    `json:"player_id"`
			PopularName string    `json:"popular_name"`
			Wickets     string    `json:"wickets"`
		} `json:"bowling"`
		Common struct {
			Batting []struct {
				BallsFaced      string `json:"balls_faced"`
				Hand            string `json:"hand"`
				ImagePath       string `json:"image_path"`
				KnownAs         string `json:"known_as"`
				Notout          string `json:"notout"`
				PlayerID        string `json:"player_id"`
				PopularName     string `json:"popular_name"`
				Position        string `json:"position"`
				PositionGroup   string `json:"position_group"`
				Runs            string `json:"runs"`
				LiveCurrentName string `json:"live_current_name,omitempty"`
			} `json:"batting"`
			Bowling []struct {
				Conceded        string `json:"conceded"`
				Hand            string `json:"hand"`
				ImagePath       string `json:"image_path"`
				KnownAs         string `json:"known_as"`
				LiveCurrentName string `json:"live_current_name,omitempty"`
				Maidens         string `json:"maidens"`
				Overs           string `json:"overs"`
				Pacespin        string `json:"pacespin"`
				PlayerID        string `json:"player_id"`
				PopularName     string `json:"popular_name"`
				Position        string `json:"position"`
				Wickets         string `json:"wickets"`
			} `json:"bowling"`
			Innings struct {
				ControlPercentage string   `json:"control_percentage"`
				DotBallPercentage string   `json:"dot_ball_percentage"`
				Event             string   `json:"event"`
				EventName         string   `json:"event_name"`
				OverLimit         string   `json:"over_limit"`
				Overs             string   `json:"overs"`
				RunRate           string   `json:"run_rate"`
				Runs              string   `json:"runs"`
				RunsSummary       []string `json:"runs_summary"`
				Target            string   `json:"target"`
				Wickets           string   `json:"wickets"`
			} `json:"innings"`
			InningsList []struct {
				Current          string `json:"current"`
				Description      string `json:"description"`
				DescriptoinShort string `json:"descriptoin_short"`
				InningsNumber    string `json:"innings_number"`
				Selected         int    `json:"selected"`
				TeamID           string `json:"team_id"`
			} `json:"innings_list"`
			InningsNumber string `json:"innings_number"`
			Match         struct {
				ControlPercentage string   `json:"control_percentage"`
				DotBallPercentage string   `json:"dot_ball_percentage"`
				ResultString      string   `json:"result_string"`
				RunsSummary       []string `json:"runs_summary"`
			} `json:"match"`
		} `json:"common"`
		Fow []struct {
			Notout int    `json:"notout"`
			Overs  string `json:"overs"`
			Player []struct {
				KnownAs     string `json:"known_as"`
				PlayerID    string `json:"player_id"`
				PopularName string `json:"popular_name"`
				Runs        string `json:"runs"`
			} `json:"player"`
			Runs string `json:"runs"`
		} `json:"fow"`
	} `json:"centre"`
	Comms []struct {
		Ball []struct {
			CommsID       string `json:"comms_id"`
			Dismissal     string `json:"dismissal"`
			Event         string `json:"event"`
			InningsNumber string `json:"innings_number"`
			IsTweet       string `json:"is_tweet"`
			OverNumber    string `json:"over_number"`
			OversActual   string `json:"overs_actual"`
			OversUnique   string `json:"overs_unique"`
			Players       string `json:"players"`
			PostText      string `json:"post_text,omitempty"`
			PreText       string `json:"pre_text,omitempty"`
			SpeedKph      string `json:"speed_kph"`
			SpeedMph      string `json:"speed_mph"`
			Text          string `json:"text"`
		} `json:"ball"`
		Batsman []struct {
			BallsFaced      string `json:"balls_faced"`
			Fours           string `json:"fours"`
			LiveCurrentName string `json:"live_current_name"`
			PlayerID        string `json:"player_id"`
			Runs            string `json:"runs"`
			Sixes           string `json:"sixes"`
		} `json:"batsman"`
		Bowler []struct {
			Conceded        string `json:"conceded"`
			LiveCurrentName string `json:"live_current_name"`
			Maidens         string `json:"maidens"`
			Overs           string `json:"overs"`
			PlayerID        string `json:"player_id"`
			Wickets         string `json:"wickets"`
		} `json:"bowler"`
		EventString    string `json:"event_string"`
		InningsNumber  string `json:"innings_number"`
		OverComplete   string `json:"over_complete"`
		OverNumber     string `json:"over_number"`
		RequiredString string `json:"required_string"`
		Runs           string `json:"runs"`
		TeamID         string `json:"team_id"`
		Wickets        string `json:"wickets"`
	} `json:"comms"`
	Description string `json:"description"`
	Innings     []struct {
		BallLimit           string      `json:"ball_limit"`
		Balls               string      `json:"balls"`
		Batted              string      `json:"batted"`
		BattingTeamID       string      `json:"batting_team_id"`
		BowlingTeamID       string      `json:"bowling_team_id"`
		Bpo                 string      `json:"bpo"`
		Byes                string      `json:"byes"`
		Event               string      `json:"event"`
		EventName           interface{} `json:"event_name"`
		Extras              string      `json:"extras"`
		InningsNumber       string      `json:"innings_number"`
		InningsNumth        string      `json:"innings_numth"`
		Lead                string      `json:"lead"`
		Legbyes             string      `json:"legbyes"`
		LiveCurrent         string      `json:"live_current"`
		LiveCurrentName     string      `json:"live_current_name"`
		Minutes             interface{} `json:"minutes"`
		Noballs             string      `json:"noballs"`
		OldPenaltyOrBonus   string      `json:"old_penalty_or_bonus"`
		OverLimit           string      `json:"over_limit"`
		OverLimitRunRate    interface{} `json:"over_limit_run_rate"`
		OverSplitLimit      string      `json:"over_split_limit"`
		Overs               string      `json:"overs"`
		OversDocked         string      `json:"overs_docked"`
		Penalties           string      `json:"penalties"`
		PenaltiesFieldEnd   string      `json:"penalties_field_end"`
		PenaltiesFieldStart string      `json:"penalties_field_start"`
		RunRate             string      `json:"run_rate"`
		Runs                string      `json:"runs"`
		Target              string      `json:"target"`
		Wickets             string      `json:"wickets"`
		Wides               string      `json:"wides"`
	} `json:"innings"`
	Live struct {
		Batting []struct {
			BallsFaced      string `json:"balls_faced"`
			BattingAverages struct {
				BattingAverage    string `json:"batting_average"`
				BattingStrikeRate string `json:"batting_strike_rate"`
				ClassCard         string `json:"class_card"`
				HighScore         string `json:"high_score"`
				Hundreds          string `json:"hundreds"`
				Innings           string `json:"innings"`
				Matches           string `json:"matches"`
				PlayerID          string `json:"player_id"`
				Runs              string `json:"runs"`
			} `json:"batting_averages"`
			BattingAveragesSeries struct {
				BattingAverage    string `json:"batting_average"`
				BattingStrikeRate string `json:"batting_strike_rate"`
				HighScore         string `json:"high_score"`
				Hundreds          string `json:"hundreds"`
				Innings           string `json:"innings"`
				Matches           string `json:"matches"`
				PlayerID          string `json:"player_id"`
				Runs              string `json:"runs"`
				SeriesTypeName    string `json:"series_type_name"`
			} `json:"batting_averages_series"`
			BattingPosition string `json:"batting_position"`
			BattingPvp      struct {
				Balls          string `json:"balls"`
				BowlerPlayerID string `json:"bowler_player_id"`
				Runs           string `json:"runs"`
			} `json:"batting_pvp"`
			BattingRecent struct {
				Balls    string `json:"balls"`
				OverSpan string `json:"over_span"`
				PlayerID string `json:"player_id"`
				Runs     string `json:"runs"`
			} `json:"batting_recent"`
			Fours           string `json:"fours"`
			InningsNumber   string `json:"innings_number"`
			LiveCurrent     string `json:"live_current"`
			LiveCurrentName string `json:"live_current_name"`
			Minutes         string `json:"minutes"`
			PlayerID        string `json:"player_id"`
			Runs            string `json:"runs"`
			Sixes           string `json:"sixes"`
			StrikeRate      string `json:"strike_rate"`
			TeamID          string `json:"team_id"`
		} `json:"batting"`
		Bowling []struct {
			BowlingAverages struct {
				Bbi            string `json:"bbi"`
				BowlingAverage string `json:"bowling_average"`
				ClassCard      string `json:"class_card"`
				EconomyRate    string `json:"economy_rate"`
				FiveWickets    string `json:"five_wickets"`
				Matches        string `json:"matches"`
				Overs          string `json:"overs"`
				PlayerID       string `json:"player_id"`
				Wickets        string `json:"wickets"`
			} `json:"bowling_averages"`
			BowlingAveragesSeries struct {
				Bbi            string `json:"bbi"`
				BowlingAverage string `json:"bowling_average"`
				EconomyRate    string `json:"economy_rate"`
				FiveWickets    string `json:"five_wickets"`
				Matches        string `json:"matches"`
				PlayerID       string `json:"player_id"`
				SeriesTypeName string `json:"series_type_name"`
				Wickets        string `json:"wickets"`
			} `json:"bowling_averages_series"`
			BowlingScoring struct {
				Dots     string `json:"dots"`
				Fours    string `json:"fours"`
				PlayerID string `json:"player_id"`
				Sixes    string `json:"sixes"`
			} `json:"bowling_scoring"`
			BowlingSpell struct {
				Balls    string `json:"balls"`
				Conceded string `json:"conceded"`
				Maidens  string `json:"maidens"`
				Overs    string `json:"overs"`
				Spell    string `json:"spell"`
				Wickets  string `json:"wickets"`
			} `json:"bowling_spell"`
			Conceded        string `json:"conceded"`
			EconomyRate     string `json:"economy_rate"`
			InningsNumber   string `json:"innings_number"`
			LiveCurrent     string `json:"live_current"`
			LiveCurrentName string `json:"live_current_name"`
			Maidens         string `json:"maidens"`
			Noballs         string `json:"noballs"`
			Overs           string `json:"overs"`
			PlayerID        string `json:"player_id"`
			TeamID          string `json:"team_id"`
			Wickets         string `json:"wickets"`
			Wides           string `json:"wides"`
		} `json:"bowling"`
		Break         string        `json:"break"`
		FieldRestrict []interface{} `json:"field_restrict"`
		Fow           []struct {
			FowOrder        string `json:"fow_order"`
			FowOvers        string `json:"fow_overs"`
			FowRuns         string `json:"fow_runs"`
			FowWickets      string `json:"fow_wickets"`
			InningsNumber   string `json:"innings_number"`
			LiveCurrent     string `json:"live_current"`
			LiveCurrentName string `json:"live_current_name"`
			OppositionID    string `json:"opposition_id"`
			OutPlayerOne    struct {
			} `json:"out_player,omitempty"`
			PartnershipOvers      string `json:"partnership_overs"`
			PartnershipRate       string `json:"partnership_rate"`
			PartnershipRuns       string `json:"partnership_runs"`
			PartnershipWicket     string `json:"partnership_wicket"`
			PartnershipWicketName string `json:"partnership_wicket_name"`
			Player                []struct {
				FowRuns         string `json:"fow_runs"`
				PartnershipRuns string `json:"partnership_runs"`
				PlayerID        string `json:"player_id"`
			} `json:"player"`
			PlayerID  string `json:"player_id"`
			TeamID    string `json:"team_id"`
			OutPlayer struct {
				BallsFaced      string `json:"balls_faced"`
				DismissalString string `json:"dismissal_string"`
				Fours           string `json:"fours"`
				Minutes         string `json:"minutes"`
				PlayerID        string `json:"player_id"`
				Runs            string `json:"runs"`
				Sixes           string `json:"sixes"`
				StrikeRate      string `json:"strike_rate"`
			} `json:"out_player,omitempty"`
		} `json:"fow"`
		Innings struct {
			BallLimit        string      `json:"ball_limit"`
			Balls            string      `json:"balls"`
			Batted           string      `json:"batted"`
			BattingTeamID    string      `json:"batting_team_id"`
			BowlingTeamID    string      `json:"bowling_team_id"`
			Bpo              string      `json:"bpo"`
			Event            string      `json:"event"`
			EventName        interface{} `json:"event_name"`
			InningsNumber    string      `json:"innings_number"`
			Lead             string      `json:"lead"`
			LiveCurrent      string      `json:"live_current"`
			LiveCurrentName  string      `json:"live_current_name"`
			OverLimit        string      `json:"over_limit"`
			OverLimitRunRate interface{} `json:"over_limit_run_rate"`
			OverSplitLimit   string      `json:"over_split_limit"`
			Overs            string      `json:"overs"`
			RemainingBalls   string      `json:"remaining_balls"`
			RemainingOvers   string      `json:"remaining_overs"`
			RemainingWickets string      `json:"remaining_wickets"`
			RequiredRunRate  interface{} `json:"required_run_rate"`
			RunRate          string      `json:"run_rate"`
			Runs             string      `json:"runs"`
			Target           string      `json:"target"`
			TeamID           string      `json:"team_id"`
			Wickets          string      `json:"wickets"`
		} `json:"innings"`
		InningsRecent []struct {
			Balls         string `json:"balls"`
			InningsNumber string `json:"innings_number"`
			OverSpan      string `json:"over_span"`
			RunRate       string `json:"run_rate"`
			Runs          string `json:"runs"`
			Wickets       string `json:"wickets"`
		} `json:"innings_recent"`
		RecentOvers [][]struct {
			Ball       string `json:"ball"`
			BallNumber string `json:"ball_number"`
			Extras     string `json:"extras"`
			OverNumber string `json:"over_number"`
		} `json:"recent_overs"`
		Review []struct {
			Remaining         string `json:"remaining"`
			RequestsPermitted string `json:"requests_permitted"`
			Successful        string `json:"successful"`
			TeamID            string `json:"team_id"`
			Unsuccessful      string `json:"unsuccessful"`
		} `json:"review"`
		Status string `json:"status"`
	} `json:"live"`
	LiveClipper struct {
	} `json:"live_clipper"`
	LiveVideo struct {
	} `json:"live_video"`
	Match struct {
		ActualDays                 string `json:"actual_days"`
		Adjusted                   string `json:"adjusted"`
		Amount                     string `json:"amount"`
		AmountBalls                string `json:"amount_balls"`
		AmountName                 string `json:"amount_name"`
		AmountType                 string `json:"amount_type"`
		AwayTeamID                 string `json:"away_team_id"`
		BallbyballSource           string `json:"ballbyball_source"`
		BattingFirstTeamID         string `json:"batting_first_team_id"`
		BitlyHash                  string `json:"bitly_hash"`
		CancelledMatch             string `json:"cancelled_match"`
		CmsMatchTitle              string `json:"cms_match_title"`
		CommentarySource           string `json:"commentary_source"`
		ContinentID                string `json:"continent_id"`
		ContinentName              string `json:"continent_name"`
		CountryAbbreviation        string `json:"country_abbreviation"`
		CountryFilename            string `json:"country_filename"`
		CountryID                  string `json:"country_id"`
		CountryName                string `json:"country_name"`
		CurrentSummary             string `json:"current_summary"`
		CurrentSummaryAbbreviation string `json:"current_summary_abbreviation"`
		Date                       string `json:"date"`
		DateString                 string `json:"date_string"`
		DaysExtended               string `json:"days_extended"`
		EndDate                    string `json:"end_date"`
		EndDateRaw                 string `json:"end_date_raw"`
		Floodlit                   string `json:"floodlit"`
		FloodlitName               string `json:"floodlit_name"`
		Followon                   string `json:"followon"`
		GeneralClassCard           string `json:"general_class_card"`
		GeneralClassID             string `json:"general_class_id"`
		GeneralClassName           string `json:"general_class_name"`
		GeneralNumber              string `json:"general_number"`
		GeneralValid               string `json:"general_valid"`
		GmtDifference              string `json:"gmt_difference"`
		GroundID                   string `json:"ground_id"`
		GroundLatitude             string `json:"ground_latitude"`
		GroundLongitude            string `json:"ground_longitude"`
		GroundName                 string `json:"ground_name"`
		GroundObjectID             string `json:"ground_object_id"`
		GroundSmallName            string `json:"ground_small_name"`
		HawkeyeSource              string `json:"hawkeye_source"`
		Head2HeadSource            string `json:"head2head_source"`
		HomeTeamID                 string `json:"home_team_id"`
		HoursString                string `json:"hours_string"`
		InternationalClassCard     string `json:"international_class_card"`
		InternationalClassID       string `json:"international_class_id"`
		InternationalClassName     string `json:"international_class_name"`
		InternationalNumber        string `json:"international_number"`
		InternationalValid         string `json:"international_valid"`
		LegacyURL                  string `json:"legacy_url"`
		LiveCommentator            string `json:"live_commentator"`
		LiveCompanion              string `json:"live_companion"`
		LiveDayNumber              string `json:"live_day_number"`
		LiveInningsNumber          string `json:"live_innings_number"`
		LiveMatch                  string `json:"live_match"`
		LiveNote                   string `json:"live_note"`
		LiveOversRemaining         string `json:"live_overs_remaining"`
		LiveOversUnique            string `json:"live_overs_unique"`
		LiveScorer                 string `json:"live_scorer"`
		LiveSessionNumber          string `json:"live_session_number"`
		LiveState                  string `json:"live_state"`
		MatchClock                 string `json:"match_clock"`
		MatchDayCountdown          string `json:"match_day_countdown"`
		MatchMinuteCountdown       string `json:"match_minute_countdown"`
		MatchPath                  string `json:"match_path"`
		MatchStatus                string `json:"match_status"`
		NeutralMatch               string `json:"neutral_match"`
		NextDatetimeGmt            string `json:"next_datetime_gmt"`
		NextDatetimeLocal          string `json:"next_datetime_local"`
		PlayerRating               string `json:"player_rating"`
		PresentDateGmt             string `json:"present_date_gmt"`
		PresentDateLocal           string `json:"present_date_local"`
		PresentDatetimeGmt         string `json:"present_datetime_gmt"`
		PresentDatetimeLocal       string `json:"present_datetime_local"`
		PresentTimeGmt             string `json:"present_time_gmt"`
		PresentTimeLocal           string `json:"present_time_local"`
		RainRule                   string `json:"rain_rule"`
		RainRuleName               string `json:"rain_rule_name"`
		RatingPromo                string `json:"rating_promo"`
		Reduced                    string `json:"reduced"`
		ReserveDaysUsed            string `json:"reserve_days_used"`
		Result                     string `json:"result"`
		ResultName                 string `json:"result_name"`
		ResultShortName            string `json:"result_short_name"`
		ScheduleNote               string `json:"schedule_note"`
		ScheduledDays              string `json:"scheduled_days"`
		ScheduledInnings           string `json:"scheduled_innings"`
		ScheduledOvers             string `json:"scheduled_overs"`
		ScorecardSource            string `json:"scorecard_source"`
		ScribbleID                 string `json:"scribble_id"`
		Season                     string `json:"season"`
		SiteID                     string `json:"site_id"`
		SiteName                   string `json:"site_name"`
		StartDate                  string `json:"start_date"`
		StartDateGmtOffset         string `json:"start_date_gmt_offset"`
		StartDateRaw               string `json:"start_date_raw"`
		StartDatetimeGmt           string `json:"start_datetime_gmt"`
		StartDatetimeGmtRaw        string `json:"start_datetime_gmt_raw"`
		StartDatetimeLocal         string `json:"start_datetime_local"`
		StartTimeGmt               string `json:"start_time_gmt"`
		StartTimeLocal             string `json:"start_time_local"`
		Team1Abbreviation          string `json:"team1_abbreviation"`
		Team1ClassID               string `json:"team1_class_id"`
		Team1CountryID             string `json:"team1_country_id"`
		Team1Filename              string `json:"team1_filename"`
		Team1ID                    string `json:"team1_id"`
		Team1LogoAltID             string `json:"team1_logo_alt_id"`
		Team1LogoEspncdn           string `json:"team1_logo_espncdn"`
		Team1LogoObjectID          string `json:"team1_logo_object_id"`
		Team1Name                  string `json:"team1_name"`
		Team1ObjectID              string `json:"team1_object_id"`
		Team1ShortName             string `json:"team1_short_name"`
		Team2Abbreviation          string `json:"team2_abbreviation"`
		Team2ClassID               string `json:"team2_class_id"`
		Team2CountryID             string `json:"team2_country_id"`
		Team2Filename              string `json:"team2_filename"`
		Team2ID                    string `json:"team2_id"`
		Team2LogoAltID             string `json:"team2_logo_alt_id"`
		Team2LogoEspncdn           string `json:"team2_logo_espncdn"`
		Team2LogoObjectID          string `json:"team2_logo_object_id"`
		Team2Name                  string `json:"team2_name"`
		Team2ObjectID              string `json:"team2_object_id"`
		Team2ShortName             string `json:"team2_short_name"`
		TiebreakerName             string `json:"tiebreaker_name"`
		TiebreakerTeamID           string `json:"tiebreaker_team_id"`
		TiebreakerType             string `json:"tiebreaker_type"`
		TimeZone                   string `json:"time_zone"`
		TossChoiceTeamID           string `json:"toss_choice_team_id"`
		TossDecision               string `json:"toss_decision"`
		TossDecisionName           string `json:"toss_decision_name"`
		TossWinnerTeamID           string `json:"toss_winner_team_id"`
		TownAka                    string `json:"town_aka"`
		TownArea                   string `json:"town_area"`
		TownID                     string `json:"town_id"`
		TownName                   string `json:"town_name"`
		TzShortName                string `json:"tz_short_name"`
		URLComponent               string `json:"url_component"`
		WatchEspnID                string `json:"watch_espn_id"`
		WeatherLocationCode        string `json:"weather_location_code"`
		WinnerTeamID               string `json:"winner_team_id"`
	} `json:"match"`
	MatchCard    string `json:"match_card"`
	MiddleColumn string `json:"middle_column"`
	Official     []struct {
		AgeDays          string `json:"age_days"`
		AgeYears         string `json:"age_years"`
		AlphaName        string `json:"alpha_name"`
		BattingHand      string `json:"batting_hand"`
		BowlingHand      string `json:"bowling_hand"`
		BowlingPacespin  string `json:"bowling_pacespin"`
		CardLong         string `json:"card_long"`
		CardQualifier    string `json:"card_qualifier"`
		CardShort        string `json:"card_short"`
		Dob              string `json:"dob"`
		KnownAs          string `json:"known_as"`
		MobileName       string `json:"mobile_name"`
		ObjectID         string `json:"object_id"`
		PlayerID         string `json:"player_id"`
		PlayerType       string `json:"player_type"`
		PlayerTypeName   string `json:"player_type_name"`
		PopularName      string `json:"popular_name"`
		PortraitAltID    string `json:"portrait_alt_id"`
		PortraitObjectID string `json:"portrait_object_id"`
		StatusID         string `json:"status_id"`
		TeamAbbreviation string `json:"team_abbreviation"`
		TeamID           string `json:"team_id"`
		TeamName         string `json:"team_name"`
		TeamShortName    string `json:"team_short_name"`
	} `json:"official"`
	OtherScores struct {
		Domestic      []interface{} `json:"domestic"`
		International []struct {
			ObjectID  string `json:"object_id"`
			StartTime string `json:"start_time"`
			Team1Desc string `json:"team1_desc"`
			Team1Name string `json:"team1_name"`
			Team2Desc string `json:"team2_desc"`
			Team2Name string `json:"team2_name"`
			URL       string `json:"url"`
		} `json:"international"`
		Others []interface{} `json:"others"`
	} `json:"other_scores"`
	ScoreSource string `json:"score_source"`
	Series      []struct {
		ClassID                string      `json:"class_id"`
		ClassName              string      `json:"class_name"`
		ContentID              string      `json:"content_id"`
		CoreRecreationID       string      `json:"core_recreation_id"`
		Date                   string      `json:"date"`
		EndDate                string      `json:"end_date"`
		EndDateRaw             string      `json:"end_date_raw"`
		FinalTypeName          string      `json:"final_type_name"`
		GroupTitle             string      `json:"group_title"`
		MajorTrophy            string      `json:"major_trophy"`
		MatchNumber            string      `json:"match_number"`
		MatchTitle             string      `json:"match_title"`
		MatchTypeName          interface{} `json:"match_type_name"`
		MultiformatPointstable string      `json:"multiformat_pointstable"`
		NumberOfHosts          string      `json:"number_of_hosts"`
		NumberOfMatches        string      `json:"number_of_matches"`
		NumberOfTeams          string      `json:"number_of_teams"`
		ObjectID               string      `json:"object_id"`
		Points                 interface{} `json:"points"`
		PrimarySeries          string      `json:"primary_series"`
		ReplayedDate           interface{} `json:"replayed_date"`
		ScheduleNote           string      `json:"schedule_note"`
		ScoreModulePosition    string      `json:"score_module_position"`
		Season                 string      `json:"season"`
		SeriesAbbreviation     interface{} `json:"series_abbreviation"`
		SeriesCategoryID       string      `json:"series_category_id"`
		SeriesCategoryName     string      `json:"series_category_name"`
		SeriesFilename         interface{} `json:"series_filename"`
		SeriesLeadAbandoned    string      `json:"series_lead_abandoned"`
		SeriesLeadCancelled    string      `json:"series_lead_cancelled"`
		SeriesLeadHowWonName   interface{} `json:"series_lead_how_won_name"`
		SeriesLeadLost         string      `json:"series_lead_lost"`
		SeriesLeadResultName   interface{} `json:"series_lead_result_name"`
		SeriesLeadTeamName     interface{} `json:"series_lead_team_name"`
		SeriesLeadTotal        string      `json:"series_lead_total"`
		SeriesLeadTypeName     interface{} `json:"series_lead_type_name"`
		SeriesLeadWon          string      `json:"series_lead_won"`
		SeriesLongDescription  string      `json:"series_long_description"`
		SeriesName             string      `json:"series_name"`
		SeriesShortName        string      `json:"series_short_name"`
		SeriesStatus           string      `json:"series_status"`
		SeriesTypeID           string      `json:"series_type_id"`
		SeriesTypeName         string      `json:"series_type_name"`
		ShortAlternateName     string      `json:"short_alternate_name"`
		SiteID                 string      `json:"site_id"`
		Slug                   string      `json:"slug"`
		StartDate              string      `json:"start_date"`
		StartDateRaw           string      `json:"start_date_raw"`
		Team1Points            interface{} `json:"team1_points"`
		Team2Points            interface{} `json:"team2_points"`
		Teams                  []struct {
			HostTeam         string      `json:"host_team"`
			ObjectID         string      `json:"object_id"`
			PrimaryTeam      string      `json:"primary_team"`
			SeriesID         string      `json:"series_id"`
			SeriesResultName interface{} `json:"series_result_name"`
			SiteID           string      `json:"site_id"`
			TeamAbbreviation string      `json:"team_abbreviation"`
			TeamFilename     string      `json:"team_filename"`
			TeamID           string      `json:"team_id"`
			TeamName         string      `json:"team_name"`
			TeamShortName    string      `json:"team_short_name"`
			URLComponent     string      `json:"url_component"`
		} `json:"teams"`
		TiebreakerID       string `json:"tiebreaker_id"`
		TiebreakerName     string `json:"tiebreaker_name"`
		TrophyAbbreviation string `json:"trophy_abbreviation"`
		TrophyClassID      string `json:"trophy_class_id"`
		TrophyCountryID    string `json:"trophy_country_id"`
		TrophyID           string `json:"trophy_id"`
		TrophyName         string `json:"trophy_name"`
		TrophyShortName    string `json:"trophy_short_name"`
		URLComponent       string `json:"url_component"`
	} `json:"series"`
	Substitute []interface{} `json:"substitute"`
	Team       []struct {
		BatsmenInSide   string `json:"batsmen_in_side"`
		ContentID       string `json:"content_id"`
		CountryID       string `json:"country_id"`
		FieldersInSide  string `json:"fielders_in_side"`
		LogoAltID       string `json:"logo_alt_id"`
		LogoEspncdn     string `json:"logo_espncdn"`
		LogoHeight      string `json:"logo_height"`
		LogoImageHeight string `json:"logo_image_height"`
		LogoImagePath   string `json:"logo_image_path"`
		LogoImageWidth  string `json:"logo_image_width"`
		LogoObjectID    string `json:"logo_object_id"`
		LogoPath        string `json:"logo_path"`
		LogoWidth       string `json:"logo_width"`
		ObjectID        string `json:"object_id"`
		Player          []struct {
			AgeDays           string `json:"age_days"`
			AgeYears          string `json:"age_years"`
			AlphaName         string `json:"alpha_name"`
			BattingHand       string `json:"batting_hand"`
			BattingStyle      string `json:"batting_style"`
			BattingStyleLong  string `json:"batting_style_long"`
			BowlingHand       string `json:"bowling_hand"`
			BowlingPacespin   string `json:"bowling_pacespin"`
			BowlingStyle      string `json:"bowling_style"`
			BowlingStyleLong  string `json:"bowling_style_long"`
			Captain           string `json:"captain"`
			CardLong          string `json:"card_long"`
			CardQualifier     string `json:"card_qualifier"`
			CardShort         string `json:"card_short"`
			Dob               string `json:"dob"`
			Keeper            string `json:"keeper"`
			KnownAs           string `json:"known_as"`
			MobileName        string `json:"mobile_name"`
			ObjectID          string `json:"object_id"`
			PlayerID          string `json:"player_id"`
			PlayerPrimaryRole string `json:"player_primary_role"`
			PlayerStyleID     string `json:"player_style_id"`
			PlayerType        string `json:"player_type"`
			PlayerTypeName    string `json:"player_type_name"`
			PopularName       string `json:"popular_name"`
			PortraitAltID     string `json:"portrait_alt_id"`
			PortraitObjectID  string `json:"portrait_object_id"`
			StatusID          string `json:"status_id"`
		} `json:"player"`
		PlayersInSide    string `json:"players_in_side"`
		SiteID           string `json:"site_id"`
		TeamAbbreviation string `json:"team_abbreviation"`
		TeamFilename     string `json:"team_filename"`
		TeamGeneralName  string `json:"team_general_name"`
		TeamID           string `json:"team_id"`
		TeamName         string `json:"team_name"`
		TeamShortName    string `json:"team_short_name"`
		URLComponent     string `json:"url_component"`
	} `json:"team"`
	Tiebreaker []interface{} `json:"tiebreaker"`
	Weather    struct {
		Forecast []interface{} `json:"forecast"`
	} `json:"weather"`
}
