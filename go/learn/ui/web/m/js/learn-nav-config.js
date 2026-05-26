(function() {
    'use strict';

    window.LAYER8M_NAV_CONFIG = {
        modules: [
            { key: 'content', label: 'Content', icon: 'content', hasSubModules: true },
            { key: 'students', label: 'Students', icon: 'students', hasSubModules: false },
            { key: 'learning', label: 'Learning', icon: 'learning', hasSubModules: false },
            { key: 'assessment', label: 'Assessment', icon: 'assessment', hasSubModules: false },
            { key: 'analytics', label: 'Analytics', icon: 'analytics', hasSubModules: false },
            { key: 'history', label: 'History', icon: 'history', hasSubModules: false },
            { key: 'collab', label: 'Collaboration', icon: 'collab', hasSubModules: false },
            { key: 'aimonitor', label: 'AI Monitor', icon: 'aimonitor', hasSubModules: false }
        ],

        content: {
            subModules: [
                { key: 'curriculum', label: 'Curriculum Framework', icon: 'content' },
                { key: 'generated', label: 'AI Generated', icon: 'aimonitor' }
            ],
            services: {
                'curriculum': [
                    { key: 'courses', label: 'Standards', icon: 'content',
                      endpoint: '/10/Course', model: 'Course', idField: 'courseId' },
                    { key: 'units', label: 'Units', icon: 'content',
                      endpoint: '/10/Unit', model: 'Unit', idField: 'unitId' },
                    { key: 'lessons', label: 'Lesson Plans', icon: 'content',
                      endpoint: '/10/Lesson', model: 'Lesson', idField: 'lessonId' },
                    { key: 'activities', label: 'Activity Bank', icon: 'content',
                      endpoint: '/10/Activity', model: 'Activity', idField: 'activityId' },
                    { key: 'worksheets', label: 'Worksheets', icon: 'content',
                      endpoint: '/10/Worksheet', model: 'Worksheet', idField: 'worksheetId' }
                ],
                'generated': [
                    { key: 'genlessons', label: 'Generated Lessons', icon: 'aimonitor',
                      endpoint: '/10/GenLesson', model: 'GeneratedLesson', idField: 'generatedLessonId' }
                ]
            }
        },

        students: {
            services: {
                'people': [
                    { key: 'students', label: 'Students', icon: 'students',
                      endpoint: '/20/Student', model: 'Student', idField: 'studentId' },
                    { key: 'profiles', label: 'Profiles', icon: 'students',
                      endpoint: '/20/Profile', model: 'StudentProfile', idField: 'profileId' },
                    { key: 'guardians', label: 'Guardians', icon: 'students',
                      endpoint: '/20/Guardian', model: 'Guardian', idField: 'guardianId' },
                    { key: 'teachers', label: 'Teachers', icon: 'students',
                      endpoint: '/20/Teacher', model: 'Teacher', idField: 'teacherId' },
                    { key: 'classrooms', label: 'Classrooms', icon: 'students',
                      endpoint: '/20/Classroom', model: 'Classroom', idField: 'classroomId' },
                    { key: 'schools', label: 'Schools', icon: 'students',
                      endpoint: '/20/School', model: 'School', idField: 'schoolId' },
                    { key: 'districts', label: 'Districts', icon: 'students',
                      endpoint: '/20/District', model: 'District', idField: 'districtId' },
                    { key: 'evals', label: 'Evaluations', icon: 'students',
                      endpoint: '/20/EvalImprt', model: 'EvalImport', idField: 'importId' }
                ]
            }
        },

        learning: {
            services: {
                'adaptive': [
                    { key: 'skills', label: 'Skills', icon: 'learning',
                      endpoint: '/30/Skill', model: 'Skill', idField: 'skillId' },
                    { key: 'mastery', label: 'Mastery', icon: 'learning',
                      endpoint: '/30/Mastery', model: 'SkillMastery', idField: 'masteryId' },
                    { key: 'paths', label: 'Paths', icon: 'learning',
                      endpoint: '/30/LearnPath', model: 'LearningPath', idField: 'pathId' },
                    { key: 'rules', label: 'Rules', icon: 'learning',
                      endpoint: '/30/AdaptRule', model: 'AdaptationRule', idField: 'ruleId' }
                ]
            }
        },

        assessment: {
            services: {
                'results': [
                    { key: 'sessions', label: 'Sessions', icon: 'assessment',
                      endpoint: '/40/LearnSess', model: 'LearningSession', idField: 'sessionId' },
                    { key: 'scores', label: 'Scores', icon: 'assessment',
                      endpoint: '/40/Score', model: 'Score', idField: 'scoreId' },
                    { key: 'benchmarks', label: 'Benchmarks', icon: 'assessment',
                      endpoint: '/40/Benchmark', model: 'Benchmark', idField: 'benchmarkId' },
                    { key: 'scans', label: 'Scans', icon: 'assessment',
                      endpoint: '/40/WkshtScan', model: 'WorksheetScan', idField: 'scanId',
                      readOnly: true }
                ]
            }
        },

        analytics: {
            services: {
                'reports': [
                    { key: 'progress', label: 'Progress', icon: 'analytics',
                      endpoint: '/50/Progress', model: 'ProgressReport', idField: 'reportId' },
                    { key: 'engagement', label: 'Engagement', icon: 'analytics',
                      endpoint: '/50/Engage', model: 'EngagementMetric', idField: 'metricId' }
                ]
            }
        },

        history: {
            services: {
                'data': [
                    { key: 'growth', label: 'Growth', icon: 'history',
                      endpoint: '/60/Growth', model: 'GrowthRecord', idField: 'recordId',
                      readOnly: true },
                    { key: 'cohorts', label: 'Cohorts', icon: 'history',
                      endpoint: '/60/Cohort', model: 'CohortSnapshot', idField: 'snapshotId',
                      readOnly: true },
                    { key: 'risk', label: 'Risk', icon: 'history',
                      endpoint: '/60/RiskAssmt', model: 'RiskAssessment', idField: 'assessmentId',
                      readOnly: true },
                    { key: 'standards', label: 'Standards', icon: 'history',
                      endpoint: '/60/StdMastry', model: 'StandardMastery', idField: 'masteryId',
                      readOnly: true },
                    { key: 'effectiveness', label: 'Effectiveness', icon: 'history',
                      endpoint: '/60/CntEffect', model: 'ContentEffect', idField: 'effectId',
                      readOnly: true }
                ]
            }
        },

        collab: {
            services: {
                'groups': [
                    { key: 'groups', label: 'Groups', icon: 'collab',
                      endpoint: '/70/Collab', model: 'CollabGroup', idField: 'groupId' },
                    { key: 'tutoring', label: 'Tutoring', icon: 'collab',
                      endpoint: '/70/TutorPair', model: 'TutorMatch', idField: 'matchId' },
                    { key: 'challenges', label: 'Challenges', icon: 'collab',
                      endpoint: '/70/Challenge', model: 'Challenge', idField: 'challengeId' }
                ]
            }
        },

        aimonitor: {
            services: {
                'data': [
                    { key: 'promptlogs', label: 'Prompt Logs', icon: 'aimonitor',
                      endpoint: '/30/PromptLog', model: 'LLMPromptLog', idField: 'logId',
                      readOnly: true },
                    { key: 'llmconfig', label: 'LLM Config', icon: 'aimonitor',
                      endpoint: '/30/LLMConfig', model: 'LLMConfig', idField: 'configId' }
                ]
            }
        },

        icons: {},
        getIcon: function(key) {
            return this.icons[key] || '';
        }
    };
})();
