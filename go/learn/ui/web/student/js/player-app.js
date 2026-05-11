/*
 * L8Learn Student Player — App Controller
 * Orchestrates the player lifecycle: welcome → activity → feedback → complete
 */
(function() {
    'use strict';

    var state = {
        currentQuestion: null,
        feedbackTimeout: null
    };

    // DOM references
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
        dom.btnStart = document.getElementById('btn-start');
        dom.btnSubmit = document.getElementById('btn-submit');
        dom.btnNext = document.getElementById('btn-next');
        dom.btnHint = document.getElementById('btn-hint');
        dom.btnDone = document.getElementById('btn-done');
        dom.welcomeGreeting = document.getElementById('welcome-greeting');
        dom.welcomeMessage = document.getElementById('welcome-message');
        dom.playerName = document.getElementById('player-name');
        dom.statPoints = document.getElementById('stat-points');
        dom.statStreak = document.getElementById('stat-streak');
        dom.completeCorrect = document.getElementById('complete-correct');
        dom.completePoints = document.getElementById('complete-points');
        dom.completeTime = document.getElementById('complete-time');
    }

    function bindEvents() {
        dom.btnStart.addEventListener('click', onStart);
        dom.btnSubmit.addEventListener('click', onSubmit);
        dom.btnNext.addEventListener('click', onNext);
        dom.btnHint.addEventListener('click', onHint);
        dom.btnDone.addEventListener('click', onDone);
    }

    async function showWelcome() {
        var name = PlayerConfig.studentName || 'Learner';
        dom.playerName.textContent = name;

        // Check if diagnostic is needed (first login)
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

        var path = await PlayerSession.startSession();
        if (!path) {
            dom.welcomeMessage.textContent = "No learning path found. Ask your teacher!";
            show(dom.welcomeScreen);
            return;
        }

        await loadActivity();
    }

    async function loadActivity() {
        var activity = await PlayerSession.loadNextActivity();
        if (!activity) {
            await finishSession();
            return;
        }

        dom.activitySkill.textContent = activity.name;
        showNextQuestion();
        show(dom.activityContainer);
    }

    function showNextQuestion() {
        var question = PlayerSession.getCurrentQuestion();
        if (!question) {
            finishSession();
            return;
        }

        state.currentQuestion = question;
        var total = PlayerSession.currentActivity.questions.length;
        dom.activityProgress.textContent = (PlayerSession.currentQuestionIndex + 1) + ' / ' + total;

        PlayerRenderer.renderQuestion(question, dom.activityContent);

        show(dom.btnSubmit);
        hide(dom.btnNext);
        if (question.hintTexts && question.hintTexts.length > 0) {
            show(dom.btnHint);
        } else {
            hide(dom.btnHint);
        }
    }

    function onSubmit() {
        var answer = PlayerRenderer.getAnswer(state.currentQuestion, dom.activityContent);
        if (!answer) return;

        var result = PlayerSession.submitAnswer(state.currentQuestion, answer);
        PlayerRenderer.showFeedback(state.currentQuestion, dom.activityContent, result.correct);
        showFeedbackOverlay(result.correct, result.points);

        hide(dom.btnSubmit);
        hide(dom.btnHint);

        // Update points display
        dom.statPoints.textContent = PlayerSession.pointsEarned + ' pts';
    }

    function onNext() {
        hide(dom.feedbackOverlay);
        if (PlayerSession.isActivityComplete()) {
            loadActivity();
        } else {
            showNextQuestion();
        }
    }

    function onHint() {
        var hint = PlayerSession.useHint(state.currentQuestion);
        if (hint) {
            var hintEl = document.createElement('div');
            hintEl.className = 'question-hint';
            hintEl.textContent = hint;
            hintEl.style.cssText = 'background:#fef3c7;padding:12px;border-radius:8px;margin-top:12px;font-size:1rem;';
            dom.activityContent.appendChild(hintEl);
        }
        if (!PlayerSession.useHint(state.currentQuestion)) {
            hide(dom.btnHint);
        }
    }

    function onDone() {
        window.location.href = '/student.html';
    }

    function showFeedbackOverlay(correct, points) {
        if (correct) {
            dom.feedbackIcon.textContent = '🎉';
            dom.feedbackMessage.textContent = 'Correct! +' + points + ' points';
            dom.feedbackMessage.style.color = 'var(--player-success)';
        } else {
            dom.feedbackIcon.textContent = '🤔';
            dom.feedbackMessage.textContent = "Not quite — let's keep going!";
            dom.feedbackMessage.style.color = 'var(--player-text)';
        }
        show(dom.feedbackOverlay);

        // Auto-dismiss after 1.5s and show Next button
        clearTimeout(state.feedbackTimeout);
        state.feedbackTimeout = setTimeout(function() {
            hide(dom.feedbackOverlay);
            show(dom.btnNext);
        }, 1500);
    }

    async function finishSession() {
        var results = await PlayerSession.completeSession();
        if (results) {
            dom.completeCorrect.textContent = results.correct + '/' + results.total;
            dom.completePoints.textContent = '+' + results.points;
            dom.completeTime.textContent = results.minutes;
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

    // Initialize when DOM is ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
