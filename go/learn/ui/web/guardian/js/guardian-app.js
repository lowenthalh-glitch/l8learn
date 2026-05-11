/*
 * L8Learn Guardian Portal — App Controller
 * Manages navigation, loads family data, renders child cards
 */
(function() {
    'use strict';

    async function init() {
        await GuardianConfig.load();
        document.getElementById('guardian-name').textContent =
            sessionStorage.getItem('userName') || 'Parent';

        bindNavigation();
        bindActivityTabs();
        loadFamily();
    }

    function bindNavigation() {
        var tabs = document.querySelectorAll('.nav-tab');
        tabs.forEach(function(tab) {
            tab.addEventListener('click', function() {
                tabs.forEach(function(t) { t.classList.remove('active'); });
                tab.classList.add('active');

                var sections = document.querySelectorAll('.portal-section');
                sections.forEach(function(s) { s.classList.remove('active'); });

                var target = tab.dataset.section;
                var section = document.querySelector('.portal-section[data-section="' + target + '"]');
                if (section) section.classList.add('active');

                onSectionChange(target);
            });
        });

        document.getElementById('btn-logout').addEventListener('click', function() {
            sessionStorage.clear();
            window.location.href = '/login.html';
        });

        var genBtn = document.getElementById('btn-generate-schedule');
        if (genBtn) {
            genBtn.addEventListener('click', generateSchedule);
        }
    }

    function bindActivityTabs() {
        var tabs = document.querySelectorAll('.tab-btn');
        tabs.forEach(function(tab) {
            tab.addEventListener('click', function() {
                tabs.forEach(function(t) { t.classList.remove('active'); });
                tab.classList.add('active');

                var contents = document.querySelectorAll('.tab-content');
                contents.forEach(function(c) { c.classList.remove('active'); });

                var target = tab.dataset.tab;
                var content = document.querySelector('.tab-content[data-tab="' + target + '"]');
                if (content) content.classList.add('active');
            });
        });
    }

    async function loadFamily() {
        var container = document.getElementById('children-grid');
        if (GuardianConfig.studentIds.length === 0) {
            container.innerHTML = '<p>No children linked to your account yet.</p>';
            return;
        }

        // Load engagement metrics for each child
        var html = '';
        for (var i = 0; i < GuardianConfig.studentIds.length; i++) {
            var studentId = GuardianConfig.studentIds[i];
            var query = 'select * from Student where studentId=' + studentId;
            var data = await GuardianConfig.get('/20/Student', query);
            if (data && data.list && data.list.length > 0) {
                var student = data.list[0];
                html += renderChildCard(student);
            }
        }
        container.innerHTML = html;

        // Load AI coach tip
        document.getElementById('coach-tip').textContent =
            "Try asking your children to explain what they learned today — " +
            "students who explain concepts retain 90% more than those who just practice.";
    }

    function renderChildCard(student) {
        var name = (student.preferredName || student.firstName) + ' ' + student.lastName;
        var initial = (student.firstName || '?')[0].toUpperCase();
        return '<div class="child-card">' +
            '<div class="child-card-header">' +
            '<div class="child-avatar">' + initial + '</div>' +
            '<div><div class="child-name">' + name + '</div>' +
            '<div class="child-grade">Grade ' + (student.gradeLevel - 2) + '</div></div>' +
            '</div>' +
            '<div class="child-progress-bar"><div class="child-progress-fill" style="width:0%"></div></div>' +
            '<div class="child-stats">Loading progress...</div>' +
            '</div>';
    }

    function onSectionChange(section) {
        switch (section) {
            case 'progress':
                loadProgress();
                break;
            case 'schedule':
                loadSchedule();
                break;
            case 'compliance':
                loadCompliance();
                break;
            case 'activities':
                // TODO: Load suggested family activities
                break;
        }
    }

    async function loadProgress() {
        // TODO: Load ProgressReport and GrowthRecord for selected child
        // Render mastery chart, growth summary, AI narrative
    }

    async function loadSchedule() {
        // TODO: Load today's DailySchedule for this family
        // Render timeline blocks
    }

    async function generateSchedule() {
        var energy = document.getElementById('parent-energy').value;
        var schedule = {
            familyId: GuardianConfig.familyId,
            scheduleDate: Math.floor(Date.now() / 1000),
            parentEnergy: energy,
            availableHours: energy === 'high' ? 5 : energy === 'medium' ? 4 : 3
        };
        await GuardianConfig.post('/30/Schedule', schedule);
        loadSchedule();
    }

    async function loadCompliance() {
        // TODO: Load StateCompliance for this family
        // Render hours bar, subjects checklist, deadlines
        var query = 'select * from StateCompliance where familyId=' + GuardianConfig.familyId;
        var data = await GuardianConfig.get('/20/Comply', query);
        if (data && data.list && data.list.length > 0) {
            var comp = data.list[0];
            var pct = comp.instructionDaysRequired > 0 ?
                Math.round(comp.instructionDaysLogged / comp.instructionDaysRequired * 100) : 0;
            var bar = document.getElementById('hours-bar');
            if (bar) bar.style.width = Math.min(pct, 100) + '%';
            var text = document.getElementById('hours-text');
            if (text) text.textContent = comp.instructionHoursLogged + ' / ' +
                comp.instructionHoursRequired + ' hours';
        }
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
