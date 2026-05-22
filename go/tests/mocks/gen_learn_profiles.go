/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Luca Lowenthal student profile — exact values from ChatGPT evaluation.
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"time"
)

func generateProfiles(store *MockDataStore) []*learn.StudentProfile {
	var list []*learn.StudentProfile
	now := time.Now().Unix()

	for i := 0; i < len(store.StudentIDs) && i < 1; i++ {
		list = append(list, generateLucaProfile(store.StudentIDs[i], now))
	}
	return list
}

func generateLucaProfile(studentId string, now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:          "PROF-0001",
		StudentId:          studentId,
		PrimaryGuardianId:  "GRD-0001",
		CreatedDate:        now,
		LastUpdated:        now,

		Metadata: &learn.ProfileMetadata{
			ProfileType:   "educational_and_developmental_working_profile",
			Purpose:       []string{"summer_learning_plan", "homeschool_support", "adaptive_learning_platform_input", "speech_and_OT_aligned_home_activities", "school_support_planning"},
			ImportantNote: "This profile is for educational planning and is not a medical diagnosis.",
			SourceBasis:   []string{"parent-provided information", "speech/language evaluation", "occupational therapy evaluation", "physical therapy evaluation", "psychomotor evaluation", "neuropsychological evaluation"},
		},

		ShortSummary: "Luca is a cheerful, energetic, social child who learns best through short, playful, movement-based, hands-on activities. He is not yet reading and needs support with speech/language, attention, fine motor skills, sensory-motor regulation, and early academic readiness.",
		LearningReadinessDescription: "Luca appears most available for learning when activities are concrete, visual, playful, adult-supported, and broken into small steps with frequent movement breaks.",
		MainStrengths: []string{
			"socially motivated", "cheerful and interactive", "strong interest in shared play",
			"physically active", "interested in books even though not reading yet",
			"enjoys movement, puzzles, Legos, art, drawing, soccer, parkour, and pool activities",
			"responds well to adult support and playful structure",
		},
		MainChallenges: []string{"attention stamina", "fine motor delays", "speech/language delays", "sensory processing differences"},
		MainLearningBarriers: []string{
			"reduced attention and impulse control during structured tasks",
			"speech/language delays affecting comprehension, expression, articulation, and narrative language",
			"fine motor and graphomotor delays affecting writing, cutting, utensil use, and school output",
			"sensory processing differences affecting attention, movement needs, oral/tactile tolerance, auditory distraction, and body awareness",
			"frustration or avoidance when tasks are too long, too verbal, too hard, or too abstract",
		},
		PrimaryGoals: []string{"improve reading readiness", "strengthen number sense", "support speech/language development", "develop fine motor skills"},

		Strengths: &learn.CategorizedStrengths{
			SocialEmotional: []string{
				"happy", "cheerful", "playful", "engaging", "interactive with adults",
				"interested in peers", "motivated by adult attention", "benefits from strong family support",
			},
			PlayAndMotivation: []string{
				"enjoys movement games", "enjoys shared play", "enjoys puzzles", "enjoys bingo",
				"enjoys Legos", "enjoys painting", "enjoys drawing", "enjoys looking at books",
				"enjoys soccer", "enjoys parkour", "enjoys pool activities",
			},
			GrossMotor: []string{
				"can ride a bike without training wheels", "can jump rope when others turn the rope",
				"can run with age-appropriate pattern", "can jump", "can hop", "can climb",
				"can kick", "can throw", "can navigate uneven surfaces",
				"physical activity is a useful learning entry point",
			},
			Academic:      []string{"interested in books", "curious about numbers"},
			Communication: []string{"verbal initiative", "social desire to communicate"},
		},

		Challenges: &learn.CategorizedChallenges{
			SpeechLanguage: []string{
				"articulation difficulties", "phonological simplification processes",
				"reduced speech intelligibility at times", "morpho-syntactic language weaknesses",
				"auditory comprehension weaknesses", "difficulty with longer or more complex directions",
				"brief and concrete discourse", "difficulty with lexical retrieval at times",
				"difficulty organizing expressive language",
			},
			AttentionExecutiveFunction: []string{
				"distractible", "impulsive", "needs frequent redirection",
				"difficulty sustaining formal structured tasks", "better performance with adult scaffolding",
				"difficulty completing longer tasks", "may rush, guess, leave the task, or lose interest",
			},
			FineMotorGraphomotor: []string{
				"difficulty writing name independently", "immature pencil grasp",
				"excessive force when writing or drawing",
				"movement often comes from shoulder/elbow rather than wrist/fingers",
				"difficulty with diagonals", "difficulty copying some forms",
				"difficulty cutting varied shapes", "fatigue with utensil use",
				"difficulty with fasteners and dressing details",
			},
			SensoryMotor: []string{
				"auditory distractibility", "tactile/oral sensitivity",
				"proprioceptive registration challenges", "vestibular challenges",
				"difficulty with tactile discrimination", "movement seeking",
				"can fatigue quickly during some motor play",
				"may rely heavily on vision for balance and body awareness",
			},
			AcademicReadiness: []string{
				"not yet reading", "needs foundational phonological awareness",
				"needs letter recognition", "needs letter-sound work",
				"needs name writing support", "needs number recognition and concrete counting practice",
				"needs classroom readiness support",
			},
		},

		Scores: &learn.WorkingScores{
			Scale:                   "0 to 10 planning score, not a clinical score",
			OverallAcademicReadiness: 4,
			ReadingReadiness:         3,
			MathReadiness:            4,
			WritingFineMotor:         3,
			SpeechLanguage:           3,
			AttentionTaskStamina:     3,
			GrossMotor:               6,
			SocialMotivation:         7,
			IndependenceDailyLiving:  4,
			ConfidenceWithLearning:   5,
		},

		LearningStyle: &learn.LearningStyle{
			PreferredModes:            []string{"visual", "kinesthetic", "hands-on"},
			BestSessionLengthMinutes:  30,
			BestActivityLengthMinutes: 6,
			MaxSeatedWorkMinutes:      7,
			BreakFrequencyMinutes:     5,
			BestTimeOfDay:             "morning",
			EffectiveActivityTypes:    []string{"movement games", "manipulatives", "object-based learning", "role play", "music and rhythm"},
			BestLearningFormula:       []string{"movement", "visual model", "hands-on task", "short speech prompt", "break", "repeat", "success"},
			WorksBestWith: []string{
				"movement games", "manipulatives", "adult modeling", "visual supports",
				"music and rhythm", "role play", "object-based learning", "short routines",
				"predictable structure", "praise for effort", "tasks with clear beginning and end",
			},
			WorksPoorlyWith: []string{
				"long verbal explanations", "long table work", "noisy environments",
				"multi-step verbal directions", "abstract worksheets", "large writing demands",
				"fast correction", "peer comparison", "open-ended finish-the-page tasks",
			},
		},

		Attention: &learn.AttentionRegulationProfile{
			MaxBookSittingTime:            "5 to 10 minutes",
			StructuredTaskStamina:         "approximately 5 to 7 minutes before break is needed",
			NeedsFrequentBreaks:           true,
			ImpulsivityPresent:            true,
			DistractibilityPresent:        true,
			FocusPreferredActivityMinutes: 15,
			FocusAcademicTaskMinutes:      6,
			LosingFocusSigns: []string{
				"leaves seat", "becomes silly", "changes topic", "rushes",
				"guesses", "avoids task", "seeks movement", "needs repeated prompts",
			},
			RegulationSupports: []string{
				"movement warm-up", "visual timer", "first/then language",
				"choice between two tasks", "short instructions", "clear endpoint",
				"immediate praise", "adult modeling", "reduce difficulty quickly", "end with easy success",
			},
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
			HighInterestActivities: []string{
				"movement", "soccer", "parkour", "pool", "books", "puzzles", "bingo",
				"Legos", "drawing", "painting", "shared games", "pretend play",
				"animal movement", "ball games",
			},
			RewardPreferences: []string{
				"verbal praise", "choice of activity", "movement game", "sticker",
				"adult attention", "displaying finished work", "small celebration", "favorite song",
			},
			AvoidAsReward:     []string{"screen time", "candy/food rewards", "large prizes"},
			AvoidedActivities: []string{"long worksheets", "writing-heavy tasks"},
		},

		Literacy: &learn.LiteracyProfile{
			CurrentLevel:      "pre-reader / emerging literacy",
			ReadingLevel:      "pre-reader",
			LetterRecognition: "some uppercase",
			PhonemicAwareness: "beginning",
			Comprehension:     "literal_only",
			SightWords:        "none yet",
			BookStaminaMinutes: 8,
			ReadingFluencyWpm:  0,
			PrimaryNeeds: []string{
				"letter recognition", "letter-sound association", "phonological awareness",
				"rhyming", "syllable awareness", "first sound identification",
				"listening comprehension", "name recognition and name writing",
			},
			PrioritySequence: []string{
				"recognize own name", "recognize letters in own name: L, U, C, A",
				"recognize uppercase letters", "match letters to sounds",
				"hear and produce rhymes", "clap syllables", "identify beginning sounds",
				"blend simple sounds orally",
				"listen to short stories and answer picture-supported questions",
				"begin CVC decoding only when earlier skills improve",
			},
			RecommendedApproach: []string{
				"no pressure to read full books yet", "daily read-aloud", "picture discussion",
				"short phonological awareness games",
				"hands-on letters using Play-Doh, magnets, blocks, sand, chalk, or tape",
				"use movement-based letter games", "use speech-aware reading instruction",
			},
			Avoid: []string{
				"long reading worksheets", "asking student to read before prerequisite skills are ready",
				"timed reading drills", "abstract phonics explanations", "large pages with too many items",
			},
		},

		Math: &learn.MathProfile{
			Level:                  "early Grade 1 readiness / Kindergarten reinforcement",
			RecommendedMode:       "hands-on, concrete, movement-based",
			Counting:              "developing",
			NumberRecognition:      "1 to 10",
			OneToOneCorrespondence: "emerging",
			PreferredTools:        []string{"blocks", "counting bears", "dice", "snacks", "Play-Doh", "sticks"},
			ErrorPatterns:         []string{"skips numbers when counting", "confuses 6 and 9"},
			PrioritySkills: []string{
				"counting objects", "one-to-one correspondence", "number recognition 1 to 10",
				"number recognition 1 to 20 when ready", "more/less", "sorting", "matching",
				"patterns", "shapes", "addition readiness", "subtraction readiness",
			},
			RecommendedActivities: []string{
				"count jumps", "count blocks", "count snacks", "number hunt",
				"build towers matching number cards", "snack patterns", "block patterns",
				"shape building with Play-Doh or sticks", "simple story problems using toys",
			},
			Avoid: []string{"long paper worksheets", "mental math without objects", "moving too quickly to symbols", "too many problems on a page"},
		},

		Speech: &learn.SpeechLanguageProfile{
			TherapyNeed:          "speech-language therapy recommended",
			ReceivesSpeechTherapy: true,
			LanguageHistory: &learn.LanguageHistory{
				LateLanguageOnset:         true,
				SingleWordsApproximateAge: "2 to 3 years",
				SentencesApproximateAge:   "around 4 years",
			},
			Articulation: &learn.ArticulationProfile{
				NotedSoundChallenges:          []string{"LL sound substitution", "trilled R / double vibrant R difficulty"},
				PhonologicalProcessesObserved: []string{"syllable omission", "sound substitution", "sound assimilation", "simplification of simple and complex word structures"},
				FunctionalImpact:             []string{"reduced intelligibility", "may affect early reading and writing acquisition", "may affect confidence when communicating"},
			},
			ExpressiveLanguage: &learn.StrengthsAndNeeds{
				Strengths: []string{"verbal initiative", "social desire to communicate", "adequate lexical base in some semantic category tasks"},
				Needs:     []string{"sentence structure", "grammar/morphosyntax", "narrative language", "word retrieval", "describing pictures", "explaining ideas", "organizing discourse"},
			},
			ReceptiveLanguage: &learn.ReceptiveLanguage{
				Needs: []string{
					"auditory comprehension", "phrase comprehension",
					"understanding verbal absurdities", "following complex instructions",
					"processing orally presented stories",
				},
				BestSupports: []string{
					"short directions", "visual supports", "gestures", "modeling",
					"one instruction at a time", "repeat and rephrase", "use objects/pictures",
				},
			},
			Clarity:      "somewhat clear",
			SpeechSounds: []string{"LL sound substitution", "trilled R difficulty"},
			CurrentGoals: []string{"improve intelligibility", "expand sentence length", "build narrative skills"},
			HelpfulPrompts: []string{"repeat slowly", "visual cues", "sentence starters"},
			RecommendedHomeGoals: []*learn.HomeGoal{
				{Goal: "Follow one-step directions consistently", Examples: []string{"Put the car in the box", "Touch the red block", "Jump to the letter M"}},
				{Goal: "Follow two-step directions with visual support", Examples: []string{"Get the ball and put it in the basket", "Pick a card and jump to the matching letter"}},
				{Goal: "Use simple complete sentences", Examples: []string{"I want ___", "I see ___", "The ___ is ___", "The animal is ___"}},
				{Goal: "Tell simple 3-part stories", Structure: []string{"first", "next", "last"}},
				{Goal: "Build vocabulary through categories", Categories: []string{"animals", "foods", "vehicles", "body parts", "actions", "colors", "places"}},
			},
		},

		FineMotor: &learn.FineMotorOTProfile{
			TherapyNeed:                 "occupational therapy support recommended",
			ReceivesOccupationalTherapy: true,
			HandDominance:               "right hand reported in psychomotor evaluation",
			PencilGrip:                  "developing",
			Tracing:                     "developing",
			Cutting:                     "emerging",
			ObservedNeeds: []string{
				"pencil grasp development", "wrist extension", "reducing excessive force during writing",
				"finger isolation and hand strength", "copying shapes and forms", "diagonal lines",
				"cutting skills", "spatial organization", "utensil use", "fasteners", "dressing details",
			},
			HelpfulWritingSupports: []string{
				"one letter per box", "start dots", "arrows", "verbal cues", "gesture cues",
				"model first, then imitate", "large writing before small writing",
				"short crayons or thick crayons", "vertical surface drawing",
				"slant board if available", "frequent breaks",
			},
			HelpfulTools: []string{"short crayons", "thick crayons", "slant board"},
			RecommendedActivities: []string{
				"Play-Doh squeezing and rolling", "clothespins", "tongs/tweezers with pom-poms",
				"stickers", "bead stringing", "pipe cleaners", "cutting strips",
				"cutting Play-Doh snakes", "large tracing paths", "mazes",
				"chalk drawing", "finger tracing in sand or shaving cream", "painting on vertical surface",
			},
			Avoid: []string{
				"long handwriting pages", "small writing before large motor preparation",
				"forcing perfect pencil grip", "making writing the first task before regulation",
				"criticizing messy output",
			},
			NameWritingStatus: &learn.NameWritingStatus{
				TargetName:            "LUCA",
				CurrentPattern:        "recognizes sequence almost completely, often LUA, and can identify missing letter when prompted",
				LettersNeedingSupport: []string{"C orientation", "U orientation", "A diagonals"},
				RecommendedPractice: []string{
					"build name with magnetic letters", "build name with Play-Doh",
					"trace name in boxes", "write one letter per box",
					"use arrows and start dots", "practice for 2 to 4 minutes only",
				},
			},
		},

		GrossMotor: &learn.GrossMotorProfile{
			OverallStatus: "gross motor skills largely age-appropriate based on PT evaluation",
			EnergyLevel:   "high",
			Strengths: []string{
				"running", "jumping", "hopping", "kicking", "throwing", "climbing",
				"navigating uneven surfaces", "bike riding without training wheels per parent report",
			},
			Needs: []string{
				"higher-level coordination", "opposite arm/leg coordination tasks",
				"balance when visual input is reduced", "motor planning in more complex tasks",
				"structured body control during movement games",
			},
			LearningUse: "Gross motor movement should be used as a learning tool before and during academic work.",
			FavoriteMovement: []string{"soccer", "parkour", "swimming", "animal walks"},
			RecommendedMovementBreaks: []string{
				"animal walks", "wall pushes", "jumping", "obstacle course", "ball toss",
				"balance beam or tape line", "crawl under chairs", "carry heavy books",
				"laundry basket push", "dance/freeze game",
			},
			MovementBreakFrequencyMinutes: 5,
		},

		Sensory: &learn.SensoryProfile{
			SensoryPattern: "mixed sensory profile with both sensitivity and registration needs",
			Sensitivities:  []string{"auditory distractors", "tactile/oral input", "some clothing textures or tags", "some food textures"},
			RegistrationNeeds: []string{"proprioceptive input", "vestibular input", "tactile discrimination", "body position awareness"},
			FunctionalImpact: []string{
				"may appear distracted or disconnected", "may seek strong movement",
				"may fatigue during postural or motor tasks",
				"may have difficulty with utensil use and tool use",
				"may struggle in noisy classroom environments",
				"may rely on vision for motor tasks",
			},
			HelpfulSupports: []string{
				"quiet work area", "reduce background noise", "heavy work before seated learning",
				"movement breaks", "visual schedule", "first/then board",
				"hands-on materials", "short tasks", "predictable routine",
			},
			Avoid: []string{
				"noisy learning environment", "long sitting without movement",
				"messy play without gradual exposure", "food pressure during academic tasks",
				"unexpected transitions",
			},
		},

		SocialEmotional: &learn.SocialEmotionalProfile{
			Confidence:      "developing",
			PeerInteraction: "enjoys peers, needs support with complex game rules",
			TurnTaking:      "emerging with adult support",
			EmotionNaming:   "can identify happy, sad, angry with picture support",
			Strengths: []string{
				"social interest", "enjoys adults", "enjoys peers",
				"pleasant and engaging", "motivated by shared play", "responds to encouragement",
			},
			Needs: []string{
				"support with frustration tolerance", "support with transitions",
				"support when language blocks communication",
				"support understanding complex game rules",
				"help expanding repetitive play themes", "help asking for help or a break",
			},
			FrustrationTriggers:  []string{"tasks too long", "tasks too hard", "consecutive errors", "language blocks communication"},
			CalmingStrategies:    []string{"deep breaths", "movement break", "switch topic"},
			RecommendedEmotionalSupports: []string{
				"calm adult voice", "predictable routine", "clear beginning and end",
				"two choices", "praise effort, not correctness only", "emotion labeling",
				"break card or break phrase", "visual schedule", "end with success",
			},
		},

		Behavior: &learn.BehaviorProfile{
			AvoidanceBehaviors: []string{"avoids long writing tasks", "leaves seat when frustrated", "rushes through to finish quickly"},
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
			Dressing: &learn.DressingProfile{
				LowerBody: "more independent", UpperBody: "can do some tasks but inconsistent",
				Socks: "difficult, needs verbal assistance", Shoes: "uses velcro; not tying laces yet",
				Zippers: "can pull zipper, difficulty starting zipper", Buttons: "not yet independent",
				BeltBuckle: "not yet independent",
			},
			Feeding: &learn.FeedingProfile{
				Spoon: "uses but fatigues; weak grasp observed", Fork: "occasional use",
				Knife: "not yet used", Preference: "often prefers eating with hands",
				FoodSelectivity: "history of texture selectivity and selective eating",
			},
			HomePracticeIdeas: []string{
				"short fork/spoon practice during snack", "button boards", "zipper-start play",
				"sock practice as a game", "visual hygiene sequence", "folding towels",
				"carrying laundry", "wiping table", "mixing batter",
			},
		},

		AdaptiveRules: &learn.AdaptiveLearningRules{
			GeneralPrinciple: "The platform should control difficulty, pacing, safety, and learning path. The AI should generate activities inside strict educational rails.",
			IfStudentIsFrustrated: []string{"pause task", "validate feeling", "offer movement break", "reduce difficulty", "give a choice", "end with easy success"},
			IfStudentIsDistracted: []string{"switch to movement-based version", "shorten the task", "use objects instead of paper", "give one-step direction", "use visual cue", "return to task after movement"},
			IfStudentIsSuccessful: []string{"repeat once for confidence", "give specific praise", "do not overextend the session", "increase difficulty only slightly", "review the skill later"},
			IfStudentAvoidsTask:   []string{"offer two choices", "reduce task to 3 items", "use first/then", "switch materials", "make it a game", "model the first response"},
			ErrorResponse:         []string{"do not say wrong first", "model the correct answer", "say it together", "try again with support", "reduce number of choices", "use visual or object prompt"},
			SuccessResponse:       []string{"praise effort specifically", "mark progress", "allow movement celebration", "keep ending positive"},
			DifficultyAdjustment: &learn.DifficultyAdjustmentTiers{
				Accuracy_0To_40:  "lower difficulty, model more, use manipulatives, shorten task",
				Accuracy_40To_70: "repeat same skill in a different playful format, add hints",
				Accuracy_70To_90: "continue practice, reduce hints, add tiny challenge",
				Accuracy_90Plus:  "mark near mastery, increase difficulty slightly, recheck later",
			},
			DefaultSessionLengthMinutes:  30,
			MaximumSessionLengthMinutes:  45,
			MaximumActivityLengthMinutes: 7,
			BreakFrequencyMinutes:        5,
			MaxConsecutiveErrors:         3,
			MaxConsecutiveCorrect:        5,
		},

		AiTutor: &learn.AITutorSettings{
			Personality: []string{"warm", "playful", "calm", "encouraging", "patient"},
			ShouldDo: []string{
				"use short sentences", "ask one question at a time", "use hints before answers",
				"use visual examples when possible", "use movement prompts", "use childs interests",
				"celebrate effort", "repeat instructions", "offer two choices",
				"redirect gently", "stop or reduce task when frustration appears",
			},
			ShouldAvoid: []string{
				"long explanations", "too many choices", "shaming mistakes", "moving too fast",
				"keeping child on screen too long", "abstract questions beyond level",
				"continuing when child is frustrated", "giving answers immediately without prompting thinking",
			},
			InstructionStyle: &learn.AIInstructionStyle{
				SentenceLength:   "short",
				QuestionsAtATime: 1,
				ChoiceCount:      2,
				Tone:             "supportive and playful",
				ReadingDemand:    "very low",
				ScreenDependency: "minimal",
			},
			ExamplePrompts: []string{
				"First we move, then we play the letter game.",
				"Do three, then break.",
				"Show me with your body.",
				fmt.Sprintf("Let%ss say it together.", ""),
				"Which one starts with /mmm/: mouse or dog?",
				"You worked hard. That was good learning.",
				"Do you want to use blocks or Play-Doh?",
			},
			PhrasesToAvoid: []string{"Sit still.", "This is easy.", "You should know this.", "Finish the whole page.", "No, wrong."},
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
				{Block: "play_based_academic_task", Minutes: 6, Purpose: "letters, sounds, numbers, book, or concept"},
				{Block: "break", Minutes: 3, Purpose: "prevent overload"},
				{Block: "speech_or_fine_motor_task", Minutes: 7, Purpose: "therapy-aligned practice"},
				{Block: "art_music_or_game", Minutes: 5, Purpose: "motivation, creativity, language"},
				{Block: "cleanup_and_celebration", Minutes: 3, Purpose: "closure and confidence"},
			},
		},

		LessonConstraints: &learn.LessonGenerationConstraints{
			MustInclude: []string{
				"movement warm-up", "visual model", "hands-on practice",
				"speech/language prompt", "fine motor or sensory-motor element",
				"short break", "positive ending",
			},
			MustAvoid: []string{
				"long worksheet-only lessons", "high reading demand",
				"long screen-based activities", "abstract verbal teaching",
				"too many items per page", "tasks longer than 7 minutes without a break",
			},
			IdealLessonLengthMinutes: 30,
			IdealWorksheetLength:     "one page maximum, large spacing, 3 to 5 sections, can be completed partly",
			WorksheetDesignRules: []string{
				"large font", "minimal clutter", "one skill per section",
				"use pictures or icons", "include tracing or coloring only briefly",
				"include parent prompts", "include movement break suggestion",
				"include success checkbox",
			},
		},

		Materials: &learn.RecommendedMaterials{
			BasicLearning: []string{
				"picture books", "alphabet cards", "magnetic letters", "number cards 1 to 20",
				"dice", "playing cards or Uno cards", "counting objects", "paper",
				"crayons", "markers", "stickers", "visual timer",
			},
			FineMotorOt: []string{
				"Play-Doh", "tongs", "tweezers", "clothespins", "pom-poms", "cotton balls",
				"beads", "pipe cleaners", "shoelace", "child-safe scissors", "glue stick",
				"short crayons", "thick crayons", "sand tray",
			},
			SpeechLanguage: []string{
				"picture cards", "family photos", "toy animals", "toy cars", "action figures",
				"simple board games", "mirror", "bubbles", "books with clear pictures",
				"household objects for I Spy",
			},
			Movement: []string{
				"ball", "beanbag", "pillows", "painters tape", "laundry basket",
				"outdoor chalk", "music speaker", "hula hoop optional", "jump rope optional",
			},
		},

		// Goals skipped — ORM bug: DomainGoals has only repeated struct fields
		// and no scalar fields, producing empty INSERT column list.

		ProgressTracking: &learn.ProgressTracking{
			TrackDaily: []string{
				"morning session completed", "afternoon session completed", "mood", "attention",
				"frustration level", "best activity", "hardest activity",
				"reading/pre-reading practice", "math practice", "speech practice",
				"fine motor practice", "movement breaks", "screen time", "parent notes",
			},
			TrackWeekly: []string{
				"skills improved", "skills still hard", "new interests",
				"avoidance patterns", "best activities", "activities to remove",
				"therapy notes", "next weeks goal",
			},
			MasteryCriteria: []string{
				"does skill independently", "does skill with minimal help",
				"does skill across 3 different days", "does skill in different activity formats",
				"can use skill in play or real life",
			},
		},

		Technology: &learn.TechnologyLimits{
			MaxScreenTimeDailyMinutes: 10,
			MaxSessionMinutes:         10,
			MaxScreenTimePercent:      20,
			AllowedUses:              []string{"learning_only"},
			ActivitiesToAvoid:        []string{"timed_tests", "competitive"},
		},
	}
}
