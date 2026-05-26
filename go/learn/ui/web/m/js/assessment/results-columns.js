(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileAssessmentResults.render;

    MobileAssessmentResults.columns = {
        LearningSession: [
            ...col.id('sessionId'),
            ...col.col('studentId', 'Student'),
            ...col.col('pathId', 'Path'),
            ...col.status('status', 'Status', null, render.sessionStatus),
            ...col.date('startTime', 'Start'),
            ...col.number('durationSeconds', 'Duration (s)'),
            ...col.number('activitiesCompleted', 'Activities'),
            ...col.number('questionsCorrect', 'Correct'),
            ...col.number('pointsEarned', 'Points')
        ],
        Score: [
            ...col.id('scoreId'),
            ...col.col('studentId', 'Student'),
            ...col.col('skillId', 'Skill'),
            ...col.col('activityId', 'Activity'),
            ...col.number('scoreValue', 'Score'),
            ...col.number('percentile', 'Percentile'),
            ...col.date('scoredDate', 'Scored'),
            ...col.number('questionsTotal', 'Total'),
            ...col.number('questionsCorrect', 'Correct')
        ],
        Benchmark: [
            ...col.id('benchmarkId'),
            ...col.col('name', 'Name'),
            ...col.enum('type', 'Type', null, render.benchmarkType),
            ...col.col('subject', 'Subject'),
            ...col.col('gradeLevel', 'Grade'),
            ...col.status('status', 'Status', null, render.contentStatus),
            ...col.number('timeLimitMinutes', 'Time Limit'),
            ...col.number('passingScore', 'Passing Score')
        ],
        WorksheetScan: [
            ...col.id('scanId'),
            ...col.col('worksheetId', 'Worksheet'),
            ...col.col('studentId', 'Student'),
            ...col.status('status', 'Status', null, render.scanStatus),
            ...col.number('correctCount', 'Correct'),
            ...col.number('totalQuestions', 'Total'),
            ...col.number('scorePercent', 'Score %'),
            ...col.number('flaggedCount', 'Flagged')
        ]
    };

    MobileAssessmentResults.columns.LearningSession[1].primary = true;
    MobileAssessmentResults.columns.LearningSession[3].secondary = true;
    MobileAssessmentResults.columns.Score[1].primary = true;
    MobileAssessmentResults.columns.Score[4].secondary = true;
    MobileAssessmentResults.columns.Benchmark[1].primary = true;
    MobileAssessmentResults.columns.Benchmark[2].secondary = true;
    MobileAssessmentResults.columns.WorksheetScan[1].primary = true;
    MobileAssessmentResults.columns.WorksheetScan[3].secondary = true;
})();
