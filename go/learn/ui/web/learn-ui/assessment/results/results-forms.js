(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = AssessmentResults.enums;

    AssessmentResults.forms = {
        LearningSession: f.form('Learning Session', [
            f.section('Session Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.reference('pathId', 'Path', 'LearningPath'),
                ...f.select('status', 'Status', enums.SESSION_STATUS),
                ...f.date('startTime', 'Start Time'),
                ...f.date('endTime', 'End Time'),
                ...f.number('durationSeconds', 'Duration (s)'),
                ...f.number('activitiesAttempted', 'Activities Attempted'),
                ...f.number('activitiesCompleted', 'Activities Completed'),
                ...f.number('questionsAnswered', 'Questions Answered'),
                ...f.number('questionsCorrect', 'Questions Correct'),
                ...f.number('hintsUsed', 'Hints Used'),
                ...f.number('pointsEarned', 'Points Earned'),
                ...f.text('deviceType', 'Device Type')
            ])
        ]),
        Score: f.form('Score', [
            f.section('Score Info', [
                ...f.reference('studentId', 'Student', 'Student', true),
                ...f.reference('skillId', 'Skill', 'Skill', true),
                ...f.reference('activityId', 'Activity', 'Activity'),
                ...f.number('scoreValue', 'Score (0-100)', true),
                ...f.number('percentile', 'Percentile'),
                ...f.date('scoredDate', 'Scored Date'),
                ...f.number('questionsTotal', 'Questions Total'),
                ...f.number('questionsCorrect', 'Questions Correct'),
                ...f.number('timeSpentSeconds', 'Time Spent (s)')
            ])
        ]),
        Benchmark: f.form('Benchmark', [
            f.section('Benchmark Info', [
                ...f.text('name', 'Name', true),
                ...f.select('type', 'Type', enums.BENCHMARK_TYPE, true),
                ...f.select('status', 'Status', enums.CONTENT_STATUS),
                ...f.number('timeLimitMinutes', 'Time Limit (min)'),
                ...f.number('passingScore', 'Passing Score (0-100)')
            ])
        ]),
        WorksheetScan: f.form('Worksheet Scan', [
            f.section('Scan Info', [
                ...f.text('scanId', 'Scan ID', false, { readOnly: true }),
                ...f.reference('worksheetId', 'Worksheet', 'Worksheet'),
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.select('status', 'Status', enums.SCAN_STATUS, false, { readOnly: true }),
                ...f.text('detectedStudentName', 'Detected Name', false, { readOnly: true }),
                ...f.number('correctCount', 'Correct', false, { readOnly: true }),
                ...f.number('totalQuestions', 'Total Questions', false, { readOnly: true }),
                ...f.number('scorePercent', 'Score %', false, { readOnly: true }),
                ...f.number('flaggedCount', 'Flagged', false, { readOnly: true })
            ])
        ])
    };
})();
