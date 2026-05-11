/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 * Licensed under the Apache License, Version 2.0
 */
var sections = {
    content: 'sections/content.html',
    students: 'sections/students.html',
    learning: 'sections/learning.html',
    assessment: 'sections/assessment.html',
    analytics: 'sections/analytics.html',
    history: 'sections/history.html',
    collab: 'sections/collab.html',
    system: 'sections/system.html'
};

var sectionInitializers = {
    content: function() { if (typeof initializeContent === 'function') initializeContent(); },
    students: function() { if (typeof initializeStudents === 'function') initializeStudents(); },
    learning: function() { if (typeof initializeLearning === 'function') initializeLearning(); },
    assessment: function() { if (typeof initializeAssessment === 'function') initializeAssessment(); },
    analytics: function() { if (typeof initializeAnalytics === 'function') initializeAnalytics(); },
    history: function() { if (typeof initializeHistory === 'function') initializeHistory(); },
    collab: function() { if (typeof initializeCollab === 'function') initializeCollab(); },
    system: function() { if (typeof initializeL8Sys === 'function') initializeL8Sys(); }
};

function updateHash(sectionName, serviceKey) {
    if (serviceKey) { window.location.hash = sectionName + '/' + serviceKey; }
    else { window.location.hash = sectionName; }
}

function getHashParts() {
    var hash = window.location.hash.replace('#', '');
    var parts = hash.split('/');
    return { section: parts[0] || '', service: parts[1] || '' };
}

document.addEventListener('click', function(e) {
    var navItem = e.target.closest('.l8-subnav-item');
    if (navItem && navItem.dataset.service) {
        var hashParts = getHashParts();
        if (hashParts.section) updateHash(hashParts.section, navItem.dataset.service);
    }
});

function loadSection(sectionName) {
    updateHash(sectionName, '');
    var contentArea = document.getElementById('content-area');
    var sectionFile = sections[sectionName];
    if (!sectionFile) {
        contentArea.innerHTML = '<div class="section-container"><h2>Section not found.</h2></div>';
        return;
    }
    contentArea.style.opacity = '0';
    contentArea.style.transform = 'translateY(20px)';
    fetch(sectionFile + '?t=' + new Date().getTime())
        .then(function(r) { if (!r.ok) throw new Error('Not found'); return r.text(); })
        .then(function(html) {
            setTimeout(function() {
                contentArea.innerHTML = html;
                var placeholder = contentArea.querySelector('[id$="-section-placeholder"]');
                if (placeholder && window.Layer8SectionGenerator) {
                    var generated = Layer8SectionGenerator.generate(sectionName);
                    var temp = document.createElement('div');
                    temp.innerHTML = generated;
                    placeholder.replaceWith.apply(placeholder, Array.from(temp.children));
                }
                setTimeout(function() {
                    contentArea.style.transition = 'opacity 0.5s ease, transform 0.5s ease';
                    contentArea.style.opacity = '1';
                    contentArea.style.transform = 'translateY(0)';
                }, 50);
                if (sectionInitializers[sectionName]) sectionInitializers[sectionName]();
            }, 200);
        })
        .catch(function() {
            contentArea.innerHTML = '<div class="section-container"><h2>Failed to load section.</h2></div>';
            contentArea.style.opacity = '1';
            contentArea.style.transform = 'translateY(0)';
        });
}
