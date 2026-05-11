/*
 * L8Learn Student Player — Session Tracker
 * Manages session lifecycle, interaction recording, and backend sync
 */
(function() {
    'use strict';

    window.PlayerSession = {
        session: null,
        path: null,
        currentActivity: null,
        currentQuestionIndex: 0,
        interactions: [],
        startTime: null,
        questionStartTime: null,
        hintsUsed: 0,
        pointsEarned: 0,

        // Start a new learning session
        async startSession() {
            this.startTime = Date.now();
            this.interactions = [];
            this.pointsEarned = 0;

            // Load the student's active learning path
            var pathQuery = 'select * from LearningPath where studentId=' +
                PlayerConfig.studentId + ' and status=1';
            var pathData = await PlayerConfig.get('/30/LearnPath', pathQuery);
            if (pathData && pathData.list && pathData.list.length > 0) {
                this.path = pathData.list[0];
            }

            // Create session record
            this.session = {
                studentId: PlayerConfig.studentId,
                pathId: this.path ? this.path.pathId : '',
                status: 1, // ACTIVE
                startTime: Math.floor(Date.now() / 1000),
                deviceType: this._detectDevice(),
                interactions: []
            };
            var result = await PlayerConfig.post('/40/LearnSess', this.session);
            if (result && result.sessionId) {
                this.session.sessionId = result.sessionId;
            }

            return this.path;
        },

        // Load the next activity from the learning path queue
        async loadNextActivity() {
            if (!this.path || !this.path.upcomingQueue || this.path.upcomingQueue.length === 0) {
                return null;
            }

            var nextStep = this.path.upcomingQueue[0];
            var activityQuery = 'select * from Activity where activityId=' + nextStep.activityId;
            var activityData = await PlayerConfig.get('/10/Activity', activityQuery);
            if (activityData && activityData.list && activityData.list.length > 0) {
                this.currentActivity = activityData.list[0];
                this.currentQuestionIndex = 0;
                this.hintsUsed = 0;
                return this.currentActivity;
            }
            return null;
        },

        // Get the current question
        getCurrentQuestion() {
            if (!this.currentActivity || !this.currentActivity.questions) {
                return null;
            }
            if (this.currentQuestionIndex >= this.currentActivity.questions.length) {
                return null;
            }
            this.questionStartTime = Date.now();
            return this.currentActivity.questions[this.currentQuestionIndex];
        },

        // Submit an answer and record the interaction
        submitAnswer(question, studentAnswer) {
            var timeSpent = Math.floor((Date.now() - this.questionStartTime) / 1000);
            var isCorrect = this._checkAnswer(question, studentAnswer);
            var points = isCorrect ? (question.points || 10) : 0;
            this.pointsEarned += points;

            var interaction = {
                activityId: this.currentActivity.activityId,
                questionId: question.questionId,
                skillId: (question.skillIds && question.skillIds.length > 0) ?
                    question.skillIds[0] : '',
                timestamp: Math.floor(Date.now() / 1000),
                result: isCorrect ? 1 : 2, // CORRECT=1, INCORRECT=2
                studentAnswer: studentAnswer,
                correctAnswer: question.correctAnswer || '',
                timeSpentSeconds: timeSpent,
                attemptNumber: 1,
                hintsUsed: this.hintsUsed,
                pointsAwarded: points,
                difficulty: this.currentActivity.difficulty
            };

            this.interactions.push(interaction);
            this.currentQuestionIndex++;
            this.hintsUsed = 0;

            return { correct: isCorrect, points: points, explanation: question.explanation };
        },

        // Use a hint for the current question
        useHint(question) {
            this.hintsUsed++;
            if (question.hintTexts && this.hintsUsed <= question.hintTexts.length) {
                return question.hintTexts[this.hintsUsed - 1];
            }
            return null;
        },

        // Check if the current activity is complete (all questions answered)
        isActivityComplete() {
            if (!this.currentActivity || !this.currentActivity.questions) {
                return true;
            }
            return this.currentQuestionIndex >= this.currentActivity.questions.length;
        },

        // Complete the session and sync to backend
        async completeSession() {
            if (!this.session) return;

            var endTime = Math.floor(Date.now() / 1000);
            var correct = this.interactions.filter(function(i) { return i.result === 1; }).length;

            this.session.status = 2; // COMPLETED
            this.session.endTime = endTime;
            this.session.durationSeconds = endTime - this.session.startTime;
            this.session.activitiesCompleted = 1;
            this.session.questionsAnswered = this.interactions.length;
            this.session.questionsCorrect = correct;
            this.session.hintsUsed = this.interactions.reduce(function(sum, i) { return sum + i.hintsUsed; }, 0);
            this.session.pointsEarned = this.pointsEarned;
            this.session.interactions = this.interactions;

            await PlayerConfig.put('/40/LearnSess', this.session);

            return {
                correct: correct,
                total: this.interactions.length,
                points: this.pointsEarned,
                minutes: Math.round(this.session.durationSeconds / 60)
            };
        },

        // Check if an answer is correct
        _checkAnswer(question, studentAnswer) {
            if (!studentAnswer) return false;

            switch (question.questionType) {
                case 1: { // SINGLE_CHOICE
                    var correctOpt = question.options.find(function(o) { return o.isCorrect; });
                    return correctOpt && correctOpt.optionId === studentAnswer;
                }
                case 2: { // MULTI_CHOICE
                    var correctIds = question.options
                        .filter(function(o) { return o.isCorrect; })
                        .map(function(o) { return o.optionId; })
                        .sort().join(',');
                    return studentAnswer.split(',').sort().join(',') === correctIds;
                }
                case 3: // NUMERIC
                case 4: { // TEXT
                    var correct = (question.correctAnswer || '').trim().toLowerCase();
                    var student = studentAnswer.trim().toLowerCase();
                    if (correct === student) return true;
                    // Try numeric equivalence (basic)
                    var cv = parseFloat(correct);
                    var sv = parseFloat(student);
                    if (!isNaN(cv) && !isNaN(sv) && Math.abs(cv - sv) < 0.001) return true;
                    return false;
                }
                default:
                    return false;
            }
        },

        _detectDevice() {
            var w = window.innerWidth;
            if (w < 768) return 'phone';
            if (w < 1024) return 'tablet';
            return 'desktop';
        }
    };
})();
