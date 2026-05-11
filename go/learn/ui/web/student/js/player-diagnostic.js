/*
 * L8Learn Student Player — Diagnostic Placement Flow
 * Runs on first login when enrollment.diagnosticComplete = false
 */
(function() {
    'use strict';

    window.PlayerDiagnostic = {
        currentSkillIndex: 0,
        currentQuestionIndex: 0,
        skills: [],
        results: [],
        questionsPerSkill: 4,

        // Check if the student needs a diagnostic
        async needsDiagnostic() {
            var query = 'select * from Enrollment where studentId=' + PlayerConfig.studentId;
            var data = await PlayerConfig.get('/20/Enroll', query);
            if (data && data.list && data.list.length > 0) {
                return !data.list[0].diagnosticComplete;
            }
            return true; // Default to needing diagnostic if no enrollment found
        },

        // Start the diagnostic flow
        async start(container) {
            this.currentSkillIndex = 0;
            this.currentQuestionIndex = 0;
            this.results = [];

            // Show welcome screen
            container.innerHTML = this._renderWelcome();
            this._bindWelcome(container);
        },

        _renderWelcome: function() {
            return '<div class="diagnostic-welcome">' +
                '<h1>Welcome!</h1>' +
                '<p>Let\'s find out what you already know!</p>' +
                '<p class="diagnostic-reassurance">This is NOT a test. There are no wrong answers.<br>' +
                'Just try your best!</p>' +
                '<p>It will take about 10-15 minutes.</p>' +
                '<button class="player-btn-start" id="diagnostic-start">Let\'s Go!</button>' +
                '</div>';
        },

        _bindWelcome: function(container) {
            var self = this;
            var btn = document.getElementById('diagnostic-start');
            if (btn) {
                btn.addEventListener('click', function() {
                    self._loadSkills(container);
                });
            }
        },

        // Load skills for the student's grade level
        async _loadSkills(container) {
            container.innerHTML = '<div class="diagnostic-loading"><p>Loading...</p></div>';

            // Load skills for the student's enrolled grade
            var query = 'select * from Skill';
            var data = await PlayerConfig.get('/30/Skill', query);
            if (data && data.list) {
                this.skills = data.list;
            }

            if (this.skills.length === 0) {
                container.innerHTML = '<div class="diagnostic-welcome">' +
                    '<h1>All Set!</h1><p>No diagnostic needed. Starting activities!</p></div>';
                setTimeout(function() { window.location.reload(); }, 2000);
                return;
            }

            // Start testing first skill
            this._testNextSkill(container);
        },

        _testNextSkill: function(container) {
            if (this.currentSkillIndex >= this.skills.length || this.currentSkillIndex >= 10) {
                this._showResults(container);
                return;
            }

            var skill = this.skills[this.currentSkillIndex];
            this.currentQuestionIndex = 0;
            this._showSkillIntro(container, skill);
        },

        _showSkillIntro: function(container, skill) {
            var self = this;
            var progress = (this.currentSkillIndex + 1) + ' / ' + Math.min(this.skills.length, 10);
            container.innerHTML = '<div class="diagnostic-skill-intro">' +
                '<div class="diagnostic-progress">Skill ' + progress + '</div>' +
                '<h2>' + (skill.name || skill.skillId) + '</h2>' +
                '<p>Let\'s try a few questions!</p>' +
                '<button class="player-btn-start" id="diagnostic-skill-start">Ready!</button>' +
                '</div>';

            document.getElementById('diagnostic-skill-start').addEventListener('click', function() {
                self._presentQuestion(container, skill);
            });
        },

        _presentQuestion: function(container, skill) {
            if (this.currentQuestionIndex >= this.questionsPerSkill) {
                // Done with this skill — evaluate and move to next
                this.currentSkillIndex++;
                this._testNextSkill(container);
                return;
            }

            var qNum = this.currentQuestionIndex + 1;
            var self = this;

            // Generate a simple diagnostic question
            // In production, this would pull from the Activity/Question pool
            container.innerHTML = '<div class="diagnostic-question">' +
                '<div class="diagnostic-progress">Question ' + qNum + ' / ' + this.questionsPerSkill + '</div>' +
                '<div class="question-prompt">Diagnostic question ' + qNum + ' for ' + (skill.name || skill.skillId) + '</div>' +
                '<div class="option-list">' +
                '<div class="option-item" data-answer="a"><div class="option-radio"></div><span>Answer A</span></div>' +
                '<div class="option-item" data-answer="b"><div class="option-radio"></div><span>Answer B</span></div>' +
                '<div class="option-item" data-answer="c"><div class="option-radio"></div><span>Answer C</span></div>' +
                '<div class="option-item" data-answer="d"><div class="option-radio"></div><span>Answer D</span></div>' +
                '</div>' +
                '<button class="player-btn-submit" id="diagnostic-submit" disabled>Check</button>' +
                '</div>';

            var options = container.querySelectorAll('.option-item');
            var submitBtn = document.getElementById('diagnostic-submit');

            options.forEach(function(opt) {
                opt.addEventListener('click', function() {
                    options.forEach(function(o) { o.classList.remove('selected'); });
                    opt.classList.add('selected');
                    submitBtn.disabled = false;
                });
            });

            submitBtn.addEventListener('click', function() {
                var selected = container.querySelector('.option-item.selected');
                if (!selected) return;

                // Record result (simulated — first option is always "correct" for diagnostic)
                var isCorrect = selected.dataset.answer === 'a';
                self.results.push({
                    skillId: skill.skillId,
                    correct: isCorrect,
                    questionIndex: self.currentQuestionIndex
                });

                // Show brief feedback
                selected.classList.add(isCorrect ? 'correct' : 'incorrect');
                setTimeout(function() {
                    self.currentQuestionIndex++;
                    self._presentQuestion(container, skill);
                }, 800);
            });
        },

        _showResults: function(container) {
            // Calculate per-skill accuracy
            var skillMap = {};
            this.results.forEach(function(r) {
                if (!skillMap[r.skillId]) skillMap[r.skillId] = { correct: 0, total: 0 };
                skillMap[r.skillId].total++;
                if (r.correct) skillMap[r.skillId].correct++;
            });

            var totalCorrect = this.results.filter(function(r) { return r.correct; }).length;
            var totalQuestions = this.results.length;

            container.innerHTML = '<div class="diagnostic-results">' +
                '<h1>All Done!</h1>' +
                '<div class="complete-stats">' +
                '<div class="complete-stat">' +
                '<span class="complete-value">' + totalCorrect + '/' + totalQuestions + '</span>' +
                '<span class="complete-label">Questions</span>' +
                '</div>' +
                '</div>' +
                '<p>Great job! We\'re building your personalized learning path...</p>' +
                '<button class="player-btn-start" id="diagnostic-done">Start Learning!</button>' +
                '</div>';

            var self = this;
            document.getElementById('diagnostic-done').addEventListener('click', function() {
                self._submitResults(skillMap);
            });
        },

        async _submitResults(skillMap) {
            // In production:
            // 1. POST SkillMastery records for each tested skill
            // 2. Update StudentProfile.readiness scores
            // 3. Trigger PATH_DECISION prompt via LLM simulator
            // 4. Create LearningPath from AI response
            // 5. Mark enrollment.diagnosticComplete = true
            // 6. Reload student player to show first activity

            // For now, just reload
            window.location.reload();
        }
    };
})();
