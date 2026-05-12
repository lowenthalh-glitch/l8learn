/*
 * L8Learn Student Player — App Controller
 * Orchestrates: welcome → lesson → step (physical/screen/worksheet) → complete
 */
(function() {
    'use strict';

    var state = {
        currentQuestion: null,
        feedbackTimeout: null,
        hintsUsed: 0
    };

    var dom = {};

    async function init() {
        await PlayerConfig.load();
        cacheDom();
        bindEvents();
        showWelcome();
    }

    function cacheDom() {
        dom.welcomeScreen = document.getElementById('welcome-screen');
        dom.activityContainer = document.getElementById('activity-container');
        dom.completeScreen = document.getElementById('complete-screen');
        dom.feedbackOverlay = document.getElementById('feedback-overlay');
        dom.feedbackIcon = document.getElementById('feedback-icon');
        dom.feedbackMessage = document.getElementById('feedback-message');
        dom.activityContent = document.getElementById('activity-content');
        dom.activitySkill = document.getElementById('activity-skill');
        dom.activityProgress = document.getElementById('activity-progress');
        dom.stepProgress = document.getElementById('step-progress');
        dom.lessonTitle = document.getElementById('lesson-title');
        dom.lessonMaterials = document.getElementById('lesson-materials');
        dom.btnStart = document.getElementById('btn-start');
        dom.btnSubmit = document.getElementById('btn-submit');
        dom.btnNext = document.getElementById('btn-next');
        dom.btnHint = document.getElementById('btn-hint');
        dom.btnDone = document.getElementById('btn-done');
        dom.btnStepDone = document.getElementById('btn-step-done');
        dom.welcomeGreeting = document.getElementById('welcome-greeting');
        dom.welcomeMessage = document.getElementById('welcome-message');
        dom.playerName = document.getElementById('player-name');
        dom.statPoints = document.getElementById('stat-points');
        dom.statStreak = document.getElementById('stat-streak');
        dom.completeCorrect = document.getElementById('complete-correct');
        dom.completePoints = document.getElementById('complete-points');
        dom.completeTime = document.getElementById('complete-time');
        dom.completeStatus = document.getElementById('complete-status');
    }

    function bindEvents() {
        dom.btnStart.addEventListener('click', onStart);
        dom.btnSubmit.addEventListener('click', onSubmit);
        dom.btnNext.addEventListener('click', onNext);
        dom.btnHint.addEventListener('click', onHint);
        dom.btnDone.addEventListener('click', onDone);
        dom.btnStepDone.addEventListener('click', onStepDone);
    }

    async function showWelcome() {
        var name = PlayerConfig.studentName || 'Learner';
        dom.playerName.textContent = name;

        // Check if diagnostic is needed
        if (window.PlayerDiagnostic) {
            var needsDiag = await PlayerDiagnostic.needsDiagnostic();
            if (needsDiag) {
                hide(dom.welcomeScreen);
                hide(dom.activityContainer);
                hide(dom.completeScreen);
                var mainEl = document.querySelector('.player-main');
                PlayerDiagnostic.start(mainEl);
                return;
            }
        }

        dom.welcomeGreeting.textContent = getGreeting() + ', ' + name + '!';
        dom.welcomeMessage.textContent = "Ready to learn something awesome today?";
        show(dom.welcomeScreen);
        hide(dom.activityContainer);
        hide(dom.completeScreen);
    }

    async function onStart() {
        hide(dom.welcomeScreen);
        dom.welcomeMessage.textContent = 'Loading your lesson...';

        var lesson = await PlayerLesson.loadLesson();
        if (!lesson) {
            dom.welcomeMessage.textContent = "No lessons ready yet. Check back soon!";
            show(dom.welcomeScreen);
            return;
        }

        showLessonHeader(lesson);
        renderCurrentStep();
        show(dom.activityContainer);
    }

    function showLessonHeader(lesson) {
        dom.lessonTitle.textContent = lesson.title || 'Today\'s Lesson';
        dom.activitySkill.textContent = lesson.topic || '';

        // Show materials if any
        if (lesson.materialsNeeded && lesson.materialsNeeded.length > 0) {
            dom.lessonMaterials.innerHTML = '<strong>Materials needed:</strong> ' +
                lesson.materialsNeeded.map(function(m) { return escapeHtml(m); }).join(', ');
            show(dom.lessonMaterials);
        } else {
            hide(dom.lessonMaterials);
        }
    }

    function renderCurrentStep() {
        var step = PlayerLesson.getCurrentStep();
        if (!step) {
            finishLesson();
            return;
        }

        dom.stepProgress.textContent = PlayerLesson.getStepProgress();
        var stepType = (step.stepType || '').toLowerCase();

        // Render the step
        PlayerRenderer.renderStep(step, dom.activityContent);

        // Show appropriate buttons based on step type
        hide(dom.btnSubmit);
        hide(dom.btnNext);
        hide(dom.btnHint);
        hide(dom.btnStepDone);

        if (stepType === 'physical' || stepType === 'hands-on' || stepType === 'worksheet') {
            show(dom.btnStepDone);
            dom.activityProgress.textContent = '';
        } else {
            // Screen step — render first question
            showNextQuestion();
        }
    }

    function showNextQuestion() {
        var question = PlayerLesson.getCurrentQuestion();
        if (!question) {
            // No more questions in this step — advance to next step
            var result = PlayerLesson.advanceStep();
            if (result === 'lesson_complete') {
                finishLesson();
            } else {
                renderCurrentStep();
            }
            return;
        }

        state.currentQuestion = question;
        state.hintsUsed = 0;
        dom.activityProgress.textContent = PlayerLesson.getQuestionProgress();

        PlayerRenderer.renderQuestion(question, dom.activityContent);

        show(dom.btnSubmit);
        hide(dom.btnNext);
        if (PlayerRenderer.hasHints(question)) {
            show(dom.btnHint);
        } else {
            hide(dom.btnHint);
        }
    }

    function onSubmit() {
        var answer = PlayerRenderer.getAnswer(state.currentQuestion, dom.activityContent);
        if (!answer) return;

        var isCorrect = PlayerLesson.checkAnswer(state.currentQuestion, answer);
        PlayerLesson.recordAnswer(isCorrect);
        var points = isCorrect ? 10 : 0;

        PlayerRenderer.showFeedback(state.currentQuestion, dom.activityContent, isCorrect);
        showFeedbackOverlay(isCorrect, points);

        hide(dom.btnSubmit);
        hide(dom.btnHint);

        dom.statPoints.textContent = (PlayerLesson.questionsCorrect * 10) + ' pts';
    }

    function onNext() {
        hide(dom.feedbackOverlay);
        var result = PlayerLesson.advanceQuestion();
        if (result === 'step_complete') {
            var stepResult = PlayerLesson.advanceStep();
            if (stepResult === 'lesson_complete') {
                finishLesson();
            } else {
                renderCurrentStep();
            }
        } else {
            showNextQuestion();
        }
    }

    function onStepDone() {
        // Physical or worksheet step marked as complete
        var stepResult = PlayerLesson.advanceStep();
        if (stepResult === 'lesson_complete') {
            finishLesson();
        } else {
            renderCurrentStep();
        }
    }

    function onHint() {
        var shown = PlayerRenderer.showHint(state.currentQuestion, state.hintsUsed, dom.activityContent);
        if (shown) {
            state.hintsUsed++;
        }
        // Check if more hints available
        var hints = state.currentQuestion.hints || state.currentQuestion.hintTexts || [];
        if (state.hintsUsed >= hints.length) {
            hide(dom.btnHint);
        }
    }

    function onDone() {
        window.location.href = '/student.html';
    }

    function showFeedbackOverlay(correct, points) {
        if (correct) {
            dom.feedbackIcon.textContent = '\uD83C\uDF89'; // party popper
            dom.feedbackMessage.textContent = 'Correct! +' + points + ' points';
            dom.feedbackMessage.style.color = 'var(--player-success)';
        } else {
            dom.feedbackIcon.textContent = '\uD83E\uDD14'; // thinking face
            dom.feedbackMessage.textContent = "Not quite \u2014 let's keep going!";
            dom.feedbackMessage.style.color = 'var(--player-text)';
        }
        show(dom.feedbackOverlay);

        clearTimeout(state.feedbackTimeout);
        state.feedbackTimeout = setTimeout(function() {
            hide(dom.feedbackOverlay);
            show(dom.btnNext);
        }, 1500);
    }

    async function finishLesson() {
        var results = await PlayerLesson.completeLesson();
        if (results) {
            dom.completeCorrect.textContent = results.correct + '/' + results.total;
            dom.completePoints.textContent = '+' + (results.correct * 10);
            dom.completeTime.textContent = results.minutes;
            if (dom.completeStatus) {
                if (results.advanced) {
                    dom.completeStatus.textContent = 'You leveled up!';
                    dom.completeStatus.style.color = 'var(--player-success)';
                } else if (results.passed) {
                    dom.completeStatus.textContent = 'Great work!';
                    dom.completeStatus.style.color = 'var(--player-primary)';
                } else {
                    dom.completeStatus.textContent = "Keep practicing \u2014 you're getting better!";
                    dom.completeStatus.style.color = 'var(--player-warning)';
                }
            }
        }
        hide(dom.activityContainer);
        show(dom.completeScreen);
    }

    function getGreeting() {
        var hour = new Date().getHours();
        if (hour < 12) return 'Good morning';
        if (hour < 17) return 'Good afternoon';
        return 'Good evening';
    }

    function show(el) { if (el) el.classList.remove('hidden'); }
    function hide(el) { if (el) el.classList.add('hidden'); }
    function escapeHtml(text) {
        var div = document.createElement('div');
        div.textContent = text || '';
        return div.innerHTML;
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
