/*
 * L8Learn Teacher Portal — Dashboard Logic
 * Loads classroom data, renders skill heatmap, risk alerts, and KPI widgets
 */
(function() {
    'use strict';

    window.TeacherDashboard = {

        async load() {
            await Promise.all([
                this.loadWidgets(),
                this.loadSkillHeatmap(),
                this.loadRiskAlerts()
            ]);
        },

        async loadWidgets() {
            // Load engagement metrics for classroom students
            var query = 'select * from EngagementMetric';
            var data = await TeacherConfig.get('/50/Engage', query);
            if (!data || !data.list) return;

            var students = data.list;
            var active = students.filter(function(s) { return s.daysSinceLastSession < 7; });
            var totalMinutes = students.reduce(function(sum, s) { return sum + (s.avgDailyMinutes || 0); }, 0);
            var avgMinutes = students.length > 0 ? Math.round(totalMinutes / students.length) : 0;

            document.getElementById('w-active-students').textContent = active.length;
            document.getElementById('w-avg-minutes').textContent = avgMinutes + ' min';

            // Load risk data
            var riskQuery = 'select * from RiskAssessment where riskLevel>1';
            var riskData = await TeacherConfig.get('/60/RiskAssmt', riskQuery);
            var atRisk = (riskData && riskData.list) ? riskData.list.length : 0;
            document.getElementById('w-at-risk').textContent = atRisk;
        },

        async loadSkillHeatmap() {
            var container = document.getElementById('skill-heatmap');

            // Load mastery data for all students in classroom
            var query = 'select * from SkillMastery';
            var data = await TeacherConfig.get('/30/Mastery', query);
            if (!data || !data.list || data.list.length === 0) {
                container.innerHTML = '<p>No mastery data yet.</p>';
                return;
            }

            // Group by student
            var byStudent = {};
            var skills = {};
            data.list.forEach(function(m) {
                if (!byStudent[m.studentId]) byStudent[m.studentId] = {};
                byStudent[m.studentId][m.skillId] = m.level;
                skills[m.skillId] = true;
            });

            var skillIds = Object.keys(skills).slice(0, 10); // Show top 10 skills

            // Build heatmap table
            var html = '<table><thead><tr><th>Student</th>';
            skillIds.forEach(function(s) {
                html += '<th>' + s.substring(0, 8) + '</th>';
            });
            html += '</tr></thead><tbody>';

            Object.keys(byStudent).forEach(function(studentId) {
                html += '<tr><td>' + studentId.substring(0, 8) + '</td>';
                skillIds.forEach(function(skillId) {
                    var level = byStudent[studentId][skillId] || 0;
                    var cls = TeacherDashboard._masteryClass(level);
                    html += '<td class="' + cls + '">' + TeacherDashboard._masteryLabel(level) + '</td>';
                });
                html += '</tr>';
            });

            html += '</tbody></table>';
            container.innerHTML = html;
        },

        async loadRiskAlerts() {
            var container = document.getElementById('risk-alerts');
            var query = 'select * from RiskAssessment where riskLevel>1';
            var data = await TeacherConfig.get('/60/RiskAssmt', query);
            if (!data || !data.list || data.list.length === 0) {
                container.innerHTML = '<p>No students at risk. Great job!</p>';
                return;
            }

            var html = '';
            data.list.forEach(function(r) {
                var badgeClass = r.riskLevel >= 3 ? 'at-risk' : 'watch';
                var label = r.riskLevel >= 3 ? 'AT RISK' : 'WATCH';
                var factors = (r.factors || []).map(function(f) { return f.description; }).join('; ');
                html += '<div class="risk-item">';
                html += '<span class="risk-badge ' + badgeClass + '">' + label + '</span>';
                html += '<div><strong>' + r.studentId.substring(0, 8) + '</strong>';
                html += '<br><small>' + factors + '</small>';
                if (r.aiRecommendation) {
                    html += '<br><em>' + r.aiRecommendation + '</em>';
                }
                html += '</div></div>';
            });

            container.innerHTML = html;
        },

        _masteryClass(level) {
            switch (level) {
                case 6: return 'mastery-exemplary';
                case 5: return 'mastery-mastered';
                case 4: return 'mastery-proficient';
                case 3: return 'mastery-developing';
                case 2: return 'mastery-emerging';
                default: return '';
            }
        },

        _masteryLabel(level) {
            switch (level) {
                case 6: return 'E';
                case 5: return 'M';
                case 4: return 'P';
                case 3: return 'D';
                case 2: return 'Em';
                default: return '-';
            }
        }
    };
})();
