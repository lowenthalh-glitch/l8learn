(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileCollabGroups.render;

    MobileCollabGroups.columns = {
        CollabGroup: [
            ...col.id('groupId'),
            ...col.col('name', 'Name'),
            ...col.enum('type', 'Type', null, render.groupType),
            ...col.status('status', 'Status', null, render.groupStatus),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.number('maxMembers', 'Max Members'),
            ...col.number('currentTeamScore', 'Team Score'),
            ...col.number('teamStreakDays', 'Streak')
        ],
        TutorMatch: [
            ...col.id('matchId'),
            ...col.col('tutorId', 'Tutor'),
            ...col.col('learnerId', 'Learner'),
            ...col.col('skillId', 'Skill'),
            ...col.number('sessionsCompleted', 'Sessions'),
            ...col.number('learnerImprovement', 'Improvement'),
            ...col.boolean('successful', 'Success'),
            ...col.date('startDate', 'Start'),
            ...col.date('endDate', 'End')
        ],
        Challenge: [
            ...col.id('challengeId'),
            ...col.col('name', 'Name'),
            ...col.enum('subject', 'Subject', null, render.subject),
            ...col.status('status', 'Status', null, render.contentStatus),
            ...col.number('teamSize', 'Team Size'),
            ...col.date('startDate', 'Start'),
            ...col.date('endDate', 'End'),
            ...col.number('targetTeamScore', 'Target Score')
        ]
    };

    MobileCollabGroups.columns.CollabGroup[1].primary = true;
    MobileCollabGroups.columns.CollabGroup[3].secondary = true;
    MobileCollabGroups.columns.TutorMatch[1].primary = true;
    MobileCollabGroups.columns.TutorMatch[6].secondary = true;
    MobileCollabGroups.columns.Challenge[1].primary = true;
    MobileCollabGroups.columns.Challenge[3].secondary = true;
})();
