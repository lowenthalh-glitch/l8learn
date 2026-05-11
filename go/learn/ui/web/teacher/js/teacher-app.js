/*
 * L8Learn Teacher Portal — App Controller
 * Handles navigation and section initialization
 */
(function() {
    'use strict';

    async function init() {
        await TeacherConfig.load();
        document.getElementById('teacher-name').textContent =
            sessionStorage.getItem('userName') || 'Teacher';

        bindNavigation();
        TeacherDashboard.load();
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
    }

    function onSectionChange(section) {
        switch (section) {
            case 'dashboard':
                TeacherDashboard.load();
                break;
            case 'students':
                // TODO: Initialize Layer8DTable for students in this classroom
                break;
            case 'worksheets':
                // TODO: Initialize Layer8DTable for worksheets by this teacher
                break;
            case 'collab':
                // TODO: Load collaboration groups, challenges, tutor pairs
                break;
            case 'scans':
                // TODO: Initialize Layer8DTable for WorksheetScans with review UI
                break;
        }
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
