/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"math/rand"
	"time"
)

func generateProfiles(store *MockDataStore) []*learn.StudentProfile {
	var list []*learn.StudentProfile
	now := time.Now().Unix()

	for i := 0; i < len(store.StudentIDs) && i < 100; i++ {
		profile := generateOneProfile(i, store.StudentIDs[i], now)
		list = append(list, profile)
	}
	return list
}

func generateOneProfile(i int, studentId string, now int64) *learn.StudentProfile {
	themes := []string{"dinosaurs", "space", "animals", "sports", "robots", "ocean", "music", "nature"}
	modes := []string{"visual", "auditory", "kinesthetic", "reading"}
	times := []string{"morning", "midday", "afternoon"}
	levels := []string{"pre-reader", "beginning", "developing", "fluent"}
	mathLevels := []string{"pre-number", "counting", "operations", "fractions"}
	personalities := []string{"warm", "playful", "calm", "encouraging", "patient"}

	theme := themes[i%len(themes)]

	return &learn.StudentProfile{
		ProfileId:   fmt.Sprintf("PROF-%04d", i+1),
		StudentId:   studentId,
		CreatedDate: now - int64(rand.Intn(90*24*3600)),
		LastUpdated: now - int64(rand.Intn(7*24*3600)),

		Metadata: &learn.ProfileMetadata{
			ProfileType:   "educational_and_developmental_working_profile",
			Purpose:       []string{"adaptive_learning_platform_input", "homeschool_support", "school_support_planning"},
			ImportantNote: "This profile is for educational planning and is not a medical diagnosis.",
			SourceBasis:   []string{"parent-provided information", "diagnostic assessment"},
		},

		ShortSummary:                fmt.Sprintf("Student %d is a cheerful learner who benefits from short, playful, movement-based activities.", i+1),
		LearningReadinessDescription: "Available for learning when activities are concrete, visual, playful, and broken into small steps.",
		MainStrengths:               []string{"socially motivated", "physically active", "enjoys " + theme},
		MainChallenges:              []string{"attention stamina", "fine motor delays"},
		MainLearningBarriers:        []string{"reduced attention during structured tasks", "fine motor delays affecting writing"},
		PrimaryGoals:                []string{"improve reading readiness", "strengthen number sense"},

		Strengths: &learn.CategorizedStrengths{
			SocialEmotional:   []string{"cheerful", "interactive", "motivated by adult attention"},
			PlayAndMotivation: []string{"enjoys " + theme, "enjoys puzzles", "enjoys drawing", "enjoys movement games"},
			GrossMotor:        []string{"can run", "can jump", "can climb", "active and energetic"},
			Academic:          []string{"interested in books", "curious about numbers"},
			Communication:     []string{"verbal initiative", "social desire to communicate"},
		},

		Challenges: &learn.CategorizedChallenges{
			SpeechLanguage:            []string{"articulation difficulties", "reduced speech intelligibility"},
			AttentionExecutiveFunction: []string{"distractible", "impulsive", "needs frequent redirection"},
			FineMotorGraphomotor:      []string{"immature pencil grasp", "difficulty with diagonals", "fatigue with writing"},
			SensoryMotor:              []string{"auditory distractibility", "movement seeking", "tactile sensitivity"},
			AcademicReadiness:         []string{"needs letter recognition", "needs number sense practice"},
		},

		Scores: &learn.WorkingScores{
			Scale:                   "0 to 10 planning score",
			OverallAcademicReadiness: int32(3 + rand.Intn(5)),
			ReadingReadiness:         int32(2 + rand.Intn(5)),
			MathReadiness:            int32(3 + rand.Intn(5)),
			WritingFineMotor:         int32(2 + rand.Intn(4)),
			SpeechLanguage:           int32(2 + rand.Intn(5)),
			AttentionTaskStamina:     int32(2 + rand.Intn(5)),
			GrossMotor:               int32(5 + rand.Intn(4)),
			SocialMotivation:         int32(5 + rand.Intn(5)),
			IndependenceDailyLiving:  int32(3 + rand.Intn(5)),
			ConfidenceWithLearning:   int32(3 + rand.Intn(5)),
		},

		LearningStyle: &learn.LearningStyle{
			PreferredModes:            []string{modes[i%len(modes)], "kinesthetic"},
			BestSessionLengthMinutes:  int32(20 + rand.Intn(15)),
			BestActivityLengthMinutes: int32(5 + rand.Intn(5)),
			MaxSeatedWorkMinutes:      int32(5 + rand.Intn(5)),
			BreakFrequencyMinutes:     int32(5 + rand.Intn(5)),
			BestTimeOfDay:             times[i%len(times)],
			EffectiveActivityTypes:    []string{"movement games", "object-based learning", "role play", "music and rhythm"},
			BestLearningFormula:       []string{"movement", "visual model", "hands-on task", "break", "repeat", "success"},
			WorksBestWith:             []string{"movement games", "manipulatives", "adult modeling", "visual supports", "short routines"},
			WorksPoorlyWith:           []string{"long verbal explanations", "long table work", "noisy environments", "abstract worksheets"},
		},

		Attention: &learn.AttentionRegulationProfile{
			MaxBookSittingTime:            "5 to 10 minutes",
			StructuredTaskStamina:         fmt.Sprintf("approximately %d to %d minutes", 5+rand.Intn(3), 7+rand.Intn(5)),
			NeedsFrequentBreaks:           true,
			ImpulsivityPresent:            i%3 == 0,
			DistractibilityPresent:        true,
			FocusPreferredActivityMinutes: int32(10 + rand.Intn(15)),
			FocusAcademicTaskMinutes:      int32(5 + rand.Intn(7)),
			LosingFocusSigns:             []string{"leaves seat", "becomes silly", "rushes", "guesses", "seeks movement"},
			RegulationSupports:           []string{"movement warm-up", "visual timer", "first/then language", "clear endpoint", "immediate praise"},
			BestInstructionStyle: &learn.InstructionStyle{
				SentenceLength:        "short",
				StepsAtATime:          1,
				UseVisualModel:        true,
				UseGesture:            true,
				RepeatDirections:      true,
				CheckUnderstanding:    true,
				AvoidLongExplanations: true,
			},
		},

		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{theme, "movement", "puzzles", "drawing", "shared games", "ball games"},
			RewardPreferences:     []string{"verbal praise", "choice of activity", "movement game", "sticker", "adult attention"},
			AvoidAsReward:         []string{"screen time", "candy/food rewards"},
			AvoidedActivities:     []string{"long worksheets", "writing-heavy tasks"},
		},

		Literacy: &learn.LiteracyProfile{
			CurrentLevel:        levels[i%len(levels)],
			ReadingLevel:        levels[i%len(levels)],
			LetterRecognition:   "some uppercase",
			PhonemicAwareness:   "beginning",
			Comprehension:       "literal_only",
			SightWords:          "none yet",
			BookStaminaMinutes:  int32(5 + rand.Intn(10)),
			ReadingFluencyWpm:   int32(0),
			PrimaryNeeds:        []string{"letter recognition", "letter-sound association", "phonological awareness", "rhyming"},
			PrioritySequence:    []string{"recognize own name", "recognize uppercase letters", "match letters to sounds", "hear and produce rhymes", "clap syllables", "identify beginning sounds"},
			RecommendedApproach: []string{"daily read-aloud", "picture discussion", "hands-on letters with Play-Doh or magnets", "movement-based letter games"},
			Avoid:               []string{"long reading worksheets", "timed reading drills", "abstract phonics explanations"},
		},

		Math: &learn.MathProfile{
			Level:                 mathLevels[i%len(mathLevels)],
			RecommendedMode:      "hands-on, concrete, movement-based",
			Counting:             "developing",
			NumberRecognition:     "1 to 10",
			OneToOneCorrespondence: "emerging",
			PreferredTools:       []string{"blocks", "counting bears", "dice", "snacks", "Play-Doh"},
			ErrorPatterns:        []string{"skips numbers when counting", "confuses 6 and 9"},
			PrioritySkills:       []string{"counting objects", "one-to-one correspondence", "number recognition 1-10", "more/less", "sorting", "patterns", "shapes"},
			RecommendedActivities: []string{"count jumps", "count blocks", "number hunt", "build towers matching number cards", "shape building with Play-Doh"},
			Avoid:                []string{"long paper worksheets", "mental math without objects", "too many problems on a page"},
		},

		Speech: &learn.SpeechLanguageProfile{
			TherapyNeed:          "speech-language therapy recommended",
			ReceivesSpeechTherapy: i%3 == 0,
			LanguageHistory: &learn.LanguageHistory{
				LateLanguageOnset:          i%2 == 0,
				SingleWordsApproximateAge: "2 to 3 years",
				SentencesApproximateAge:   "around 4 years",
			},
			Articulation: &learn.ArticulationProfile{
				NotedSoundChallenges:          []string{"R sounds", "L blends"},
				PhonologicalProcessesObserved: []string{"syllable omission", "sound substitution"},
				FunctionalImpact:             []string{"reduced intelligibility", "may affect early reading"},
			},
			ExpressiveLanguage: &learn.StrengthsAndNeeds{
				Strengths: []string{"verbal initiative", "social desire to communicate"},
				Needs:     []string{"sentence structure", "grammar", "narrative language", "word retrieval"},
			},
			ReceptiveLanguage: &learn.ReceptiveLanguage{
				Needs:        []string{"auditory comprehension", "following complex instructions"},
				BestSupports: []string{"short directions", "visual supports", "gestures", "one instruction at a time"},
			},
			Clarity:        "somewhat_clear",
			SpeechSounds:   []string{"LL substitution", "trilled R difficulty"},
			CurrentGoals:   []string{"improve intelligibility", "expand sentence length", "build narrative skills"},
			HelpfulPrompts: []string{"repeat slowly", "visual cues", "sentence starters"},
			RecommendedHomeGoals: []*learn.HomeGoal{
				{Goal: "Follow one-step directions consistently", Examples: []string{"Put the car in the box", "Touch the red block"}},
				{Goal: "Use simple complete sentences", Examples: []string{"I want ___", "I see ___", "The ___ is ___"}},
				{Goal: "Tell simple 3-part stories", Structure: []string{"first", "next", "last"}},
			},
		},

		FineMotor: &learn.FineMotorOTProfile{
			TherapyNeed:              "occupational therapy support recommended",
			ReceivesOccupationalTherapy: i%4 == 0,
			HandDominance:            "right",
			PencilGrip:               "developing",
			Tracing:                  "developing",
			Cutting:                  "emerging",
			ObservedNeeds:            []string{"pencil grasp development", "finger isolation", "copying shapes", "cutting skills"},
			HelpfulWritingSupports:   []string{"one letter per box", "start dots", "arrows", "model first then imitate", "large writing before small"},
			HelpfulTools:             []string{"short crayons", "thick crayons", "slant board"},
			RecommendedActivities:    []string{"Play-Doh squeezing", "clothespins", "tongs with pom-poms", "stickers", "bead stringing", "chalk drawing"},
			Avoid:                    []string{"long handwriting pages", "forcing perfect pencil grip", "making writing first task"},
			NameWritingStatus: &learn.NameWritingStatus{
				TargetName:             "STUDENT " + fmt.Sprintf("%d", i+1),
				CurrentPattern:         "recognizes sequence, sometimes omits letters",
				LettersNeedingSupport:  []string{"diagonal letters"},
				RecommendedPractice:    []string{"build name with magnetic letters", "trace in boxes", "2 to 4 minutes only"},
			},
		},

		GrossMotor: &learn.GrossMotorProfile{
			OverallStatus:                "largely age-appropriate",
			EnergyLevel:                  "high",
			Strengths:                    []string{"running", "jumping", "climbing", "kicking", "throwing"},
			Needs:                        []string{"higher-level coordination", "balance with reduced visual input"},
			LearningUse:                  "Gross motor movement should be used as a learning tool before and during academic work.",
			FavoriteMovement:             []string{"soccer", "parkour", "swimming", "animal walks"},
			RecommendedMovementBreaks:    []string{"animal walks", "wall pushes", "jumping", "obstacle course", "ball toss", "dance/freeze"},
			MovementBreakFrequencyMinutes: 5,
		},

		Sensory: &learn.SensoryProfile{
			SensoryPattern:    "mixed sensory profile with both sensitivity and registration needs",
			Sensitivities:     []string{"auditory distractors", "tactile/oral input", "some clothing textures"},
			RegistrationNeeds: []string{"proprioceptive input", "vestibular input", "body position awareness"},
			FunctionalImpact:  []string{"may appear distracted", "may seek strong movement", "may fatigue during motor tasks"},
			HelpfulSupports:   []string{"quiet work area", "heavy work before seated learning", "movement breaks", "visual schedule", "short tasks"},
			Avoid:             []string{"noisy learning environment", "long sitting without movement", "unexpected transitions"},
		},

		SocialEmotional: &learn.SocialEmotionalProfile{
			Confidence:                   "developing",
			PeerInteraction:              "interested but needs support with game rules",
			TurnTaking:                   "emerging with adult support",
			EmotionNaming:                "can identify happy, sad, angry with picture support",
			Strengths:                    []string{"social interest", "enjoys adults", "pleasant and engaging", "responds to encouragement"},
			Needs:                        []string{"frustration tolerance", "transitions support", "help asking for a break"},
			FrustrationTriggers:          []string{"tasks too long", "tasks too hard", "consecutive errors"},
			CalmingStrategies:            []string{"deep breaths", "movement break", "switch topic"},
			RecommendedEmotionalSupports: []string{"calm adult voice", "predictable routine", "two choices", "praise effort not correctness", "break card", "end with success"},
		},

		Behavior: &learn.BehaviorProfile{
			AvoidanceBehaviors:  []string{"avoids long writing tasks", "leaves seat when frustrated", "rushes through to finish quickly"},
			RedirectStrategies: []string{"offer movement break", "switch to hands-on version", "reduce task size", "use first/then"},
			SuccessfulSupports: []string{"adult modeling", "visual timer", "choice between two tasks", "end with easy success"},
		},

		Health: &learn.HealthSafety{
			MedicalConditions: []string{"none reported"},
			VisionConcerns:    "no concerns reported",
			HearingConcerns:   "no concerns reported",
			SafetyConcerns:    []string{"impulsive near roads", "needs supervision near water"},
		},

		DailyLiving: &learn.DailyLivingProfile{
			Toileting: &learn.ToiletingProfile{Urination: "independent", BowelHygiene: "requires assistance"},
			Hygiene:   &learn.HygieneProfile{HandWashing: "independent", ToothBrushing: "independent with supervision", Bathing: "requires help"},
			Dressing:  &learn.DressingProfile{LowerBody: "more independent", UpperBody: "inconsistent", Socks: "difficult", Shoes: "velcro", Zippers: "can pull, difficulty starting", Buttons: "not yet independent", BeltBuckle: "not yet independent"},
			Feeding:   &learn.FeedingProfile{Spoon: "uses but fatigues", Fork: "occasional use", Knife: "not yet used", Preference: "often prefers hands", FoodSelectivity: "history of texture selectivity"},
			HomePracticeIdeas: []string{"short fork/spoon practice during snack", "button boards", "zipper-start play", "sock practice as game", "visual hygiene sequence"},
		},

		AdaptiveRules: &learn.AdaptiveLearningRules{
			GeneralPrinciple:       "The platform should control difficulty, pacing, safety, and learning path.",
			IfStudentIsFrustrated:  []string{"pause task", "validate feeling", "offer movement break", "reduce difficulty", "give a choice", "end with easy success"},
			IfStudentIsDistracted:  []string{"switch to movement-based version", "shorten the task", "use objects instead of paper", "give one-step direction"},
			IfStudentIsSuccessful:  []string{"repeat once for confidence", "give specific praise", "do not overextend", "increase difficulty only slightly"},
			IfStudentAvoidsTask:    []string{"offer two choices", "reduce task to 3 items", "use first/then", "make it a game", "model the first response"},
			ErrorResponse:          []string{"do not say wrong first", "model correct answer", "say it together", "try again with support"},
			SuccessResponse:        []string{"praise effort specifically", "mark progress", "allow movement celebration"},
			DifficultyAdjustment: &learn.DifficultyAdjustmentTiers{
				Accuracy_0To_40:  "lower difficulty, model more, use manipulatives, shorten task",
				Accuracy_40To_70: "repeat same skill in different playful format, add hints",
				Accuracy_70To_90: "continue practice, reduce hints, add tiny challenge",
				Accuracy_90Plus:  "mark near mastery, increase difficulty slightly, recheck later",
			},
			DefaultSessionLengthMinutes: 30,
			MaximumSessionLengthMinutes: 45,
			MaximumActivityLengthMinutes: 7,
			BreakFrequencyMinutes:       5,
			MaxConsecutiveErrors:        3,
			MaxConsecutiveCorrect:       5,
		},

		AiTutor: &learn.AITutorSettings{
			Personality: []string{personalities[i%len(personalities)], personalities[(i+1)%len(personalities)]},
			ShouldDo:    []string{"use short sentences", "ask one question at a time", "use hints before answers", "use childs interests", "celebrate effort", "redirect gently"},
			ShouldAvoid: []string{"long explanations", "shaming mistakes", "moving too fast", "continuing when frustrated", "giving answers immediately"},
			InstructionStyle: &learn.AIInstructionStyle{
				SentenceLength:   "short",
				QuestionsAtATime: 1,
				ChoiceCount:      2,
				Tone:             "supportive and playful",
				ReadingDemand:    "very low",
				ScreenDependency: "minimal",
			},
			ExamplePrompts:  []string{"First we move, then we play the letter game.", "Do three, then break.", "Show me with your body.", "Which one starts with /mmm/: mouse or dog?"},
			PhrasesToAvoid:  []string{"Sit still.", "This is easy.", "You should know this.", "Finish the whole page.", "No, wrong."},
		},

		ProgramSettings: &learn.ProgramSettings{
			SessionsPerDay:            2,
			SessionLengthMinutes:      30,
			DaysPerWeek:               5,
			OptionalLightReviewDays:   []string{"Saturday", "Sunday"},
			MaxSeatedWorkMinutes:      7,
			MovementBreakEveryMinutes: 5,
			DailySessionStructure: []*learn.SessionBlock{
				{Block: "movement_warmup", Minutes: 5, Purpose: "regulate attention and body"},
				{Block: "play_based_academic_task", Minutes: 6, Purpose: "letters, sounds, numbers, or concept"},
				{Block: "break", Minutes: 3, Purpose: "prevent overload"},
				{Block: "speech_or_fine_motor_task", Minutes: 7, Purpose: "therapy-aligned practice"},
				{Block: "art_music_or_game", Minutes: 5, Purpose: "motivation, creativity, language"},
				{Block: "cleanup_and_celebration", Minutes: 3, Purpose: "closure and confidence"},
			},
		},

		LessonConstraints: &learn.LessonGenerationConstraints{
			MustInclude:             []string{"movement warm-up", "visual model", "hands-on practice", "speech/language prompt", "short break", "positive ending"},
			MustAvoid:               []string{"long worksheet-only lessons", "high reading demand", "long screen-based activities", "abstract verbal teaching", "tasks longer than 7 minutes without break"},
			IdealLessonLengthMinutes: 30,
			IdealWorksheetLength:    "one page maximum, large spacing, 3 to 5 sections",
			WorksheetDesignRules:    []string{"large font", "minimal clutter", "one skill per section", "use pictures or icons", "include parent prompts", "include success checkbox"},
		},

		Materials: &learn.RecommendedMaterials{
			BasicLearning:  []string{"picture books", "alphabet cards", "magnetic letters", "number cards 1-20", "dice", "counting objects", "crayons", "stickers", "visual timer"},
			FineMotorOt:    []string{"Play-Doh", "tongs", "tweezers", "clothespins", "pom-poms", "beads", "pipe cleaners", "child-safe scissors", "short crayons", "sand tray"},
			SpeechLanguage: []string{"picture cards", "toy animals", "action figures", "simple board games", "mirror", "bubbles", "books with clear pictures"},
			Movement:       []string{"ball", "beanbag", "pillows", "painters tape", "outdoor chalk", "music speaker"},
		},

		// Goals skipped — ORM bug: DomainGoals has only repeated struct fields
		// and no scalar fields, producing empty INSERT column list.

		ProgressTracking: &learn.ProgressTracking{
			TrackDaily:      []string{"session completed", "mood", "attention", "frustration level", "best activity", "screen time"},
			TrackWeekly:     []string{"skills improved", "skills still hard", "new interests", "avoidance patterns", "next weeks goal"},
			MasteryCriteria: []string{"does skill independently", "does skill with minimal help", "does skill across 3 different days", "does skill in different formats"},
		},

		Technology: &learn.TechnologyLimits{
			MaxScreenTimeDailyMinutes: int32(10 + rand.Intn(20)),
			MaxSessionMinutes:         int32(10),
			MaxScreenTimePercent:       20,
			AllowedUses:               []string{"learning_only"},
			ActivitiesToAvoid:         []string{"timed_tests", "competitive"},
		},
	}
}
