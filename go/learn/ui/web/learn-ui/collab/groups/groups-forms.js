(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = CollabGroups.enums;

    CollabGroups.forms = {
        CollabGroup: f.form('Collaboration Group', [
            f.section('Group Info', [
                ...f.text('name', 'Name', true),
                ...f.select('type', 'Type', enums.GROUP_TYPE, true),
                ...f.select('status', 'Status', enums.GROUP_STATUS),
                ...f.textarea('description', 'Description'),
                ...f.number('maxMembers', 'Max Members'),
                ...f.text('classroomId', 'Classroom'),
                ...f.select('subject', 'Subject', enums.SUBJECT),
                ...f.number('teamScoreGoal', 'Team Score Goal'),
                ...f.date('deadline', 'Deadline')
            ])
        ]),
        TutorMatch: f.form('Tutor Match', [
            f.section('Match Info', [
                ...f.reference('tutorId', 'Tutor', 'Student', true),
                ...f.reference('learnerId', 'Learner', 'Student', true),
                ...f.reference('skillId', 'Skill', 'Skill', true),
                ...f.reference('groupId', 'Group', 'CollabGroup'),
                ...f.select('tutorMastery', 'Tutor Mastery', enums.MASTERY_LEVEL),
                ...f.select('learnerStart', 'Learner Start', enums.MASTERY_LEVEL),
                ...f.select('learnerEnd', 'Learner End', enums.MASTERY_LEVEL),
                ...f.number('sessionsCompleted', 'Sessions Completed'),
                ...f.number('learnerImprovement', 'Improvement'),
                ...f.checkbox('successful', 'Successful'),
                ...f.number('tutorPointsEarned', 'Tutor Points'),
                ...f.number('learnerPointsEarned', 'Learner Points'),
                ...f.date('startDate', 'Start Date'),
                ...f.date('endDate', 'End Date')
            ])
        ]),
        Challenge: f.form('Challenge', [
            f.section('Challenge Info', [
                ...f.text('name', 'Name', true),
                ...f.textarea('description', 'Description'),
                ...f.text('classroomId', 'Classroom'),
                ...f.select('subject', 'Subject', enums.SUBJECT),
                ...f.select('status', 'Status', enums.CONTENT_STATUS),
                ...f.number('teamSize', 'Team Size'),
                ...f.date('startDate', 'Start Date'),
                ...f.date('endDate', 'End Date'),
                ...f.number('dailyRequirement', 'Daily Requirement'),
                ...f.checkbox('weakestMemberBonus', 'Weakest Member Bonus'),
                ...f.number('targetTeamScore', 'Target Team Score')
            ])
        ])
    };
})();
