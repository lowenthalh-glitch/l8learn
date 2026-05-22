/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
StudentProfile form — all sections matching the full ChatGPT evaluation JSON.
*/
(function() {
    'use strict';
    var f = window.Layer8FormFactory;

    StudentsPeople.forms.StudentProfile = f.form('Student Profile', [
        f.section('Summary', [
            ...f.reference('studentId', 'Student', 'Student'),
            ...f.textarea('shortSummary', 'Summary'),
            ...f.text('learningReadinessDescription', 'Learning Readiness'),
            ...f.text('mainStrengths', 'Main Strengths'),
            ...f.text('mainChallenges', 'Main Challenges'),
            ...f.text('mainLearningBarriers', 'Main Learning Barriers'),
            ...f.text('primaryGoals', 'Primary Goals')
        ]),
        f.section('Metadata', [
            ...f.text('metadata.profileType', 'Profile Type'),
            ...f.text('metadata.purpose', 'Purpose'),
            ...f.text('metadata.importantNote', 'Important Note'),
            ...f.text('metadata.sourceBasis', 'Source Basis')
        ]),
        f.section('Readiness Scores', [
            ...f.number('scores.overallAcademicReadiness', 'Academic Readiness'),
            ...f.number('scores.readingReadiness', 'Reading Readiness'),
            ...f.number('scores.mathReadiness', 'Math Readiness'),
            ...f.number('scores.writingFineMotor', 'Writing/Fine Motor'),
            ...f.number('scores.speechLanguage', 'Speech/Language'),
            ...f.number('scores.attentionTaskStamina', 'Attention Stamina'),
            ...f.number('scores.grossMotor', 'Gross Motor'),
            ...f.number('scores.socialMotivation', 'Social Motivation'),
            ...f.number('scores.independenceDailyLiving', 'Independence'),
            ...f.number('scores.confidenceWithLearning', 'Confidence')
        ]),
        f.section('Strengths', [
            ...f.text('strengths.socialEmotional', 'Social-Emotional'),
            ...f.text('strengths.playAndMotivation', 'Play & Motivation'),
            ...f.text('strengths.grossMotor', 'Gross Motor'),
            ...f.text('strengths.academic', 'Academic'),
            ...f.text('strengths.communication', 'Communication')
        ]),
        f.section('Challenges', [
            ...f.text('challenges.speechLanguage', 'Speech/Language'),
            ...f.text('challenges.attentionExecutiveFunction', 'Attention/Executive Function'),
            ...f.text('challenges.fineMotorGraphomotor', 'Fine Motor/Graphomotor'),
            ...f.text('challenges.sensoryMotor', 'Sensory-Motor'),
            ...f.text('challenges.academicReadiness', 'Academic Readiness')
        ]),
        f.section('Learning Style', [
            ...f.text('learningStyle.preferredModes', 'Preferred Modes'),
            ...f.number('learningStyle.bestSessionLengthMinutes', 'Best Session (min)'),
            ...f.number('learningStyle.bestActivityLengthMinutes', 'Best Activity (min)'),
            ...f.number('learningStyle.maxSeatedWorkMinutes', 'Max Seated Work (min)'),
            ...f.number('learningStyle.breakFrequencyMinutes', 'Break Frequency (min)'),
            ...f.text('learningStyle.bestTimeOfDay', 'Best Time of Day'),
            ...f.text('learningStyle.bestLearningFormula', 'Best Learning Formula'),
            ...f.text('learningStyle.worksBestWith', 'Works Best With'),
            ...f.text('learningStyle.worksPoorlyWith', 'Works Poorly With')
        ]),
        f.section('Attention & Regulation', [
            ...f.text('attention.maxBookSittingTime', 'Max Book Sitting Time'),
            ...f.text('attention.structuredTaskStamina', 'Structured Task Stamina'),
            ...f.checkbox('attention.needsFrequentBreaks', 'Needs Frequent Breaks'),
            ...f.checkbox('attention.impulsivityPresent', 'Impulsivity Present'),
            ...f.checkbox('attention.distractibilityPresent', 'Distractibility Present'),
            ...f.number('attention.focusPreferredActivityMinutes', 'Focus Preferred Activity (min)'),
            ...f.number('attention.focusAcademicTaskMinutes', 'Focus Academic Task (min)'),
            ...f.text('attention.losingFocusSigns', 'Losing Focus Signs'),
            ...f.text('attention.regulationSupports', 'Regulation Supports')
        ]),
        f.section('Motivation', [
            ...f.text('motivation.highInterestActivities', 'High Interest Activities'),
            ...f.text('motivation.rewardPreferences', 'Reward Preferences'),
            ...f.text('motivation.avoidAsReward', 'Avoid as Reward'),
            ...f.text('motivation.avoidedActivities', 'Avoided Activities')
        ]),
        f.section('Literacy / Pre-Reading', [
            ...f.text('literacy.currentLevel', 'Current Level'),
            ...f.text('literacy.readingLevel', 'Reading Level'),
            ...f.text('literacy.letterRecognition', 'Letter Recognition'),
            ...f.text('literacy.phonemicAwareness', 'Phonemic Awareness'),
            ...f.text('literacy.sightWords', 'Sight Words'),
            ...f.text('literacy.comprehension', 'Comprehension'),
            ...f.number('literacy.bookStaminaMinutes', 'Book Stamina (min)'),
            ...f.number('literacy.readingFluencyWpm', 'Fluency (WPM)'),
            ...f.text('literacy.primaryNeeds', 'Primary Needs'),
            ...f.text('literacy.prioritySequence', 'Priority Sequence'),
            ...f.text('literacy.recommendedApproach', 'Recommended Approach'),
            ...f.text('literacy.avoid', 'Avoid')
        ]),
        f.section('Math', [
            ...f.text('math.level', 'Level'),
            ...f.text('math.recommendedMode', 'Recommended Mode'),
            ...f.text('math.counting', 'Counting'),
            ...f.text('math.numberRecognition', 'Number Recognition'),
            ...f.text('math.oneToOneCorrespondence', 'One-to-One Correspondence'),
            ...f.text('math.prioritySkills', 'Priority Skills'),
            ...f.text('math.recommendedActivities', 'Recommended Activities'),
            ...f.text('math.errorPatterns', 'Error Patterns'),
            ...f.text('math.avoid', 'Avoid')
        ]),
        f.section('Speech & Language', [
            ...f.text('speech.therapyNeed', 'Therapy Need'),
            ...f.checkbox('speech.receivesSpeechTherapy', 'Receives Speech Therapy'),
            ...f.text('speech.clarity', 'Clarity'),
            ...f.text('speech.speechSounds', 'Speech Sounds'),
            ...f.text('speech.currentGoals', 'Current Goals'),
            ...f.text('speech.helpfulPrompts', 'Helpful Prompts'),
            ...f.text('speech.articulation.notedSoundChallenges', 'Articulation: Sound Challenges'),
            ...f.text('speech.articulation.phonologicalProcessesObserved', 'Articulation: Phonological Processes'),
            ...f.text('speech.articulation.functionalImpact', 'Articulation: Functional Impact'),
            ...f.text('speech.expressiveLanguage.strengths', 'Expressive: Strengths'),
            ...f.text('speech.expressiveLanguage.needs', 'Expressive: Needs'),
            ...f.text('speech.receptiveLanguage.needs', 'Receptive: Needs'),
            ...f.text('speech.receptiveLanguage.bestSupports', 'Receptive: Best Supports'),
            ...f.text('speech.languageHistory.singleWordsApproximateAge', 'First Words Age'),
            ...f.text('speech.languageHistory.sentencesApproximateAge', 'First Sentences Age'),
            ...f.checkbox('speech.languageHistory.lateLanguageOnset', 'Late Language Onset')
        ]),
        f.section('Fine Motor / OT', [
            ...f.text('fineMotor.therapyNeed', 'Therapy Need'),
            ...f.checkbox('fineMotor.receivesOccupationalTherapy', 'Receives OT'),
            ...f.text('fineMotor.handDominance', 'Hand Dominance'),
            ...f.text('fineMotor.pencilGrip', 'Pencil Grip'),
            ...f.text('fineMotor.tracing', 'Tracing'),
            ...f.text('fineMotor.cutting', 'Cutting'),
            ...f.text('fineMotor.observedNeeds', 'Observed Needs'),
            ...f.text('fineMotor.helpfulWritingSupports', 'Helpful Writing Supports'),
            ...f.text('fineMotor.helpfulTools', 'Helpful Tools'),
            ...f.text('fineMotor.recommendedActivities', 'Recommended Activities'),
            ...f.text('fineMotor.avoid', 'Avoid'),
            ...f.text('fineMotor.nameWritingStatus.targetName', 'Name Writing: Target'),
            ...f.text('fineMotor.nameWritingStatus.currentPattern', 'Name Writing: Current Pattern'),
            ...f.text('fineMotor.nameWritingStatus.lettersNeedingSupport', 'Name Writing: Letters Needing Support'),
            ...f.text('fineMotor.nameWritingStatus.recommendedPractice', 'Name Writing: Practice')
        ]),
        f.section('Gross Motor', [
            ...f.text('grossMotor.overallStatus', 'Overall Status'),
            ...f.text('grossMotor.energyLevel', 'Energy Level'),
            ...f.text('grossMotor.strengths', 'Strengths'),
            ...f.text('grossMotor.needs', 'Needs'),
            ...f.text('grossMotor.learningUse', 'Learning Use'),
            ...f.text('grossMotor.favoriteMovement', 'Favorite Movement'),
            ...f.text('grossMotor.recommendedMovementBreaks', 'Recommended Movement Breaks'),
            ...f.number('grossMotor.movementBreakFrequencyMinutes', 'Break Frequency (min)')
        ]),
        f.section('Sensory Profile', [
            ...f.text('sensory.sensoryPattern', 'Sensory Pattern'),
            ...f.text('sensory.sensitivities', 'Sensitivities'),
            ...f.text('sensory.registrationNeeds', 'Registration Needs'),
            ...f.text('sensory.functionalImpact', 'Functional Impact'),
            ...f.text('sensory.helpfulSupports', 'Helpful Supports'),
            ...f.text('sensory.avoid', 'Avoid')
        ]),
        f.section('Social-Emotional', [
            ...f.text('socialEmotional.confidence', 'Confidence'),
            ...f.text('socialEmotional.peerInteraction', 'Peer Interaction'),
            ...f.text('socialEmotional.turnTaking', 'Turn Taking'),
            ...f.text('socialEmotional.emotionNaming', 'Emotion Naming'),
            ...f.text('socialEmotional.strengths', 'Strengths'),
            ...f.text('socialEmotional.needs', 'Needs'),
            ...f.text('socialEmotional.frustrationTriggers', 'Frustration Triggers'),
            ...f.text('socialEmotional.calmingStrategies', 'Calming Strategies'),
            ...f.text('socialEmotional.recommendedEmotionalSupports', 'Recommended Supports')
        ]),
        f.section('Behavior', [
            ...f.text('behavior.avoidanceBehaviors', 'Avoidance Behaviors'),
            ...f.text('behavior.redirectStrategies', 'Redirect Strategies'),
            ...f.text('behavior.successfulSupports', 'Successful Supports')
        ]),
        f.section('Daily Living', [
            ...f.text('dailyLiving.toileting.urination', 'Toileting: Urination'),
            ...f.text('dailyLiving.toileting.bowelHygiene', 'Toileting: Bowel Hygiene'),
            ...f.text('dailyLiving.hygiene.handWashing', 'Hygiene: Hand Washing'),
            ...f.text('dailyLiving.hygiene.toothBrushing', 'Hygiene: Tooth Brushing'),
            ...f.text('dailyLiving.hygiene.bathing', 'Hygiene: Bathing'),
            ...f.text('dailyLiving.dressing.lowerBody', 'Dressing: Lower Body'),
            ...f.text('dailyLiving.dressing.upperBody', 'Dressing: Upper Body'),
            ...f.text('dailyLiving.dressing.socks', 'Dressing: Socks'),
            ...f.text('dailyLiving.dressing.shoes', 'Dressing: Shoes'),
            ...f.text('dailyLiving.dressing.zippers', 'Dressing: Zippers'),
            ...f.text('dailyLiving.dressing.buttons', 'Dressing: Buttons'),
            ...f.text('dailyLiving.feeding.spoon', 'Feeding: Spoon'),
            ...f.text('dailyLiving.feeding.fork', 'Feeding: Fork'),
            ...f.text('dailyLiving.feeding.preference', 'Feeding: Preference'),
            ...f.text('dailyLiving.feeding.foodSelectivity', 'Feeding: Food Selectivity'),
            ...f.text('dailyLiving.homePracticeIdeas', 'Home Practice Ideas')
        ]),
        f.section('Health & Safety', [
            ...f.text('health.medicalConditions', 'Medical Conditions'),
            ...f.text('health.visionConcerns', 'Vision Concerns'),
            ...f.text('health.hearingConcerns', 'Hearing Concerns'),
            ...f.text('health.safetyConcerns', 'Safety Concerns')
        ]),
        f.section('Adaptive Learning Rules', [
            ...f.textarea('adaptiveRules.generalPrinciple', 'General Principle'),
            ...f.text('adaptiveRules.ifStudentIsFrustrated', 'If Frustrated'),
            ...f.text('adaptiveRules.ifStudentIsDistracted', 'If Distracted'),
            ...f.text('adaptiveRules.ifStudentIsSuccessful', 'If Successful'),
            ...f.text('adaptiveRules.ifStudentAvoidsTask', 'If Avoids Task'),
            ...f.text('adaptiveRules.errorResponse', 'Error Response'),
            ...f.text('adaptiveRules.successResponse', 'Success Response'),
            ...f.number('adaptiveRules.defaultSessionLengthMinutes', 'Default Session (min)'),
            ...f.number('adaptiveRules.maximumActivityLengthMinutes', 'Max Activity (min)'),
            ...f.number('adaptiveRules.breakFrequencyMinutes', 'Break Frequency (min)'),
            ...f.number('adaptiveRules.maxConsecutiveErrors', 'Max Consecutive Errors'),
            ...f.number('adaptiveRules.maxConsecutiveCorrect', 'Max Consecutive Correct')
        ]),
        f.section('AI Tutor Settings', [
            ...f.text('aiTutor.personality', 'Personality'),
            ...f.text('aiTutor.shouldDo', 'Should Do'),
            ...f.text('aiTutor.shouldAvoid', 'Should Avoid'),
            ...f.text('aiTutor.instructionStyle.sentenceLength', 'Sentence Length'),
            ...f.number('aiTutor.instructionStyle.questionsAtATime', 'Questions at a Time'),
            ...f.number('aiTutor.instructionStyle.choiceCount', 'Choice Count'),
            ...f.text('aiTutor.instructionStyle.tone', 'Tone'),
            ...f.text('aiTutor.instructionStyle.readingDemand', 'Reading Demand'),
            ...f.text('aiTutor.instructionStyle.screenDependency', 'Screen Dependency'),
            ...f.text('aiTutor.examplePrompts', 'Example Prompts'),
            ...f.text('aiTutor.phrasesToAvoid', 'Phrases to Avoid')
        ]),
        f.section('Program Settings', [
            ...f.number('programSettings.sessionsPerDay', 'Sessions Per Day'),
            ...f.number('programSettings.sessionLengthMinutes', 'Session Length (min)'),
            ...f.number('programSettings.daysPerWeek', 'Days Per Week'),
            ...f.text('programSettings.optionalLightReviewDays', 'Optional Review Days'),
            ...f.number('programSettings.maxSeatedWorkMinutes', 'Max Seated Work (min)'),
            ...f.number('programSettings.movementBreakEveryMinutes', 'Movement Break Every (min)')
        ]),
        f.section('Lesson Constraints', [
            ...f.text('lessonConstraints.mustInclude', 'Must Include'),
            ...f.text('lessonConstraints.mustAvoid', 'Must Avoid'),
            ...f.number('lessonConstraints.idealLessonLengthMinutes', 'Ideal Lesson (min)'),
            ...f.text('lessonConstraints.idealWorksheetLength', 'Ideal Worksheet Length'),
            ...f.text('lessonConstraints.worksheetDesignRules', 'Worksheet Design Rules')
        ]),
        f.section('Recommended Materials', [
            ...f.text('materials.basicLearning', 'Basic Learning'),
            ...f.text('materials.fineMotorOt', 'Fine Motor / OT'),
            ...f.text('materials.speechLanguage', 'Speech/Language'),
            ...f.text('materials.movement', 'Movement')
        ]),
        f.section('Technology Limits', [
            ...f.number('technology.maxScreenTimeDailyMinutes', 'Max Screen Time (min/day)'),
            ...f.number('technology.maxSessionMinutes', 'Max Session (min)'),
            ...f.number('technology.maxScreenTimePercent', 'Max Screen Time (%)'),
            ...f.text('technology.allowedUses', 'Allowed Uses'),
            ...f.text('technology.activitiesToAvoid', 'Activities to Avoid')
        ]),
        f.section('Progress Tracking', [
            ...f.text('progressTracking.trackDaily', 'Track Daily'),
            ...f.text('progressTracking.trackWeekly', 'Track Weekly'),
            ...f.text('progressTracking.masteryCriteria', 'Mastery Criteria')
        ])
    ]);
})();
