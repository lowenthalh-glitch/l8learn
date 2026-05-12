/*
 * L8Learn Student Player — Lesson Manager
 * Fetches and navigates AI-generated lessons (GeneratedLesson)
 */
(function() {
    'use strict';

    window.PlayerLesson = {
        lesson: null,
        currentStepIndex: 0,
        currentQuestionIndex: 0,
        questionsCorrect: 0,
        questionsTotal: 0,
        lessonStartTime: null,

        // Load the next ready GeneratedLesson for this student
        async loadLesson() {
            var query = 'select * from GeneratedLesson where studentId=' +
                PlayerConfig.studentId + ' and status=2 limit 1 page 0';
            var data = await PlayerConfig.get('/10/GenLesson', query);
            if (data && data.list && data.list.length > 0) {
                this.lesson = data.list[0];
                this.currentStepIndex = 0;
                this.currentQuestionIndex = 0;
                this.questionsCorrect = 0;
                this.questionsTotal = 0;
                this.lessonStartTime = Date.now();
                // Mark as IN_PROGRESS
                await PlayerConfig.put('/10/GenLesson', {
                    generatedLessonId: this.lesson.generatedLessonId,
                    status: 3, // IN_PROGRESS
                    startedAt: Math.floor(Date.now() / 1000)
                });
                return this.lesson;
            }
            return null;
        },

        // Get the current step
        getCurrentStep() {
            if (!this.lesson || !this.lesson.steps) return null;
            if (this.currentStepIndex >= this.lesson.steps.length) return null;
            return this.lesson.steps[this.currentStepIndex];
        },

        // Get the current question within the current screen step
        getCurrentQuestion() {
            var step = this.getCurrentStep();
            if (!step || !step.questions) return null;
            if (this.currentQuestionIndex >= step.questions.length) return null;
            return step.questions[this.currentQuestionIndex];
        },

        // Advance to next question or next step
        advanceQuestion() {
            var step = this.getCurrentStep();
            if (step && step.questions && this.currentQuestionIndex < step.questions.length - 1) {
                this.currentQuestionIndex++;
                return 'next_question';
            }
            return 'step_complete';
        },

        // Advance to the next step
        advanceStep() {
            this.currentStepIndex++;
            this.currentQuestionIndex = 0;
            if (this.currentStepIndex >= this.lesson.steps.length) {
                return 'lesson_complete';
            }
            return 'next_step';
        },

        // Record a question result
        recordAnswer(correct) {
            this.questionsTotal++;
            if (correct) this.questionsCorrect++;
        },

        // Check answer for a GeneratedQuestion
        checkAnswer(question, studentAnswer) {
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
                    var cv = parseFloat(correct);
                    var sv = parseFloat(student);
                    if (!isNaN(cv) && !isNaN(sv) && Math.abs(cv - sv) < 0.001) return true;
                    return false;
                }
                default:
                    return false;
            }
        },

        // Complete the lesson: PUT results and request next lesson
        async completeLesson() {
            if (!this.lesson) return null;
            var actualMinutes = Math.round((Date.now() - this.lessonStartTime) / 60000);
            var passed = this.questionsCorrect >= (this.lesson.minCorrectToPass || 0);

            await PlayerConfig.put('/10/GenLesson', {
                generatedLessonId: this.lesson.generatedLessonId,
                status: 4, // COMPLETED
                questionsCorrect: this.questionsCorrect,
                questionsTotal: this.questionsTotal,
                actualMinutes: actualMinutes,
                completedAt: Math.floor(Date.now() / 1000)
            });

            var result = {
                correct: this.questionsCorrect,
                total: this.questionsTotal,
                minutes: actualMinutes,
                passed: passed,
                advanced: this.questionsCorrect >= (this.lesson.minCorrectToAdvance || 0)
            };

            // Request next lesson generation (fire and forget)
            this._requestNextLesson();

            this.lesson = null;
            return result;
        },

        // Ask the backend to generate the next lesson
        async _requestNextLesson() {
            try {
                await PlayerConfig.post('/10/GenLesson', {
                    studentId: PlayerConfig.studentId,
                    status: 1 // GENERATING
                });
            } catch (e) {
                // Non-critical — next lesson will be generated when student returns
            }
        },

        // Get step progress display text
        getStepProgress() {
            if (!this.lesson || !this.lesson.steps) return '';
            return 'Step ' + (this.currentStepIndex + 1) + ' of ' + this.lesson.steps.length;
        },

        // Get question progress within current step
        getQuestionProgress() {
            var step = this.getCurrentStep();
            if (!step || !step.questions) return '';
            return 'Question ' + (this.currentQuestionIndex + 1) + ' of ' + step.questions.length;
        }
    };
})();
