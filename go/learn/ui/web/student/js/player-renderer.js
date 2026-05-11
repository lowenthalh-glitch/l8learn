/*
 * L8Learn Student Player — Activity Renderer
 * Renders questions by type into the activity container
 */
(function() {
    'use strict';

    window.PlayerRenderer = {

        // Render a question into the activity-content container
        renderQuestion(question, container) {
            container.innerHTML = '';

            const prompt = document.createElement('div');
            prompt.className = 'question-prompt';
            prompt.textContent = question.promptText;
            container.appendChild(prompt);

            if (question.promptMediaUrl) {
                const img = document.createElement('img');
                img.src = question.promptMediaUrl;
                img.className = 'question-media';
                img.alt = 'Question image';
                container.appendChild(img);
            }

            switch (question.questionType) {
                case 1: // SINGLE_CHOICE
                    this._renderSingleChoice(question, container);
                    break;
                case 2: // MULTI_CHOICE
                    this._renderMultiChoice(question, container);
                    break;
                case 3: // NUMERIC
                    this._renderNumeric(question, container);
                    break;
                case 4: // TEXT
                    this._renderText(question, container);
                    break;
                case 5: // DRAG_DROP
                    this._renderDragDrop(question, container);
                    break;
                case 6: // DRAWING
                    this._renderDrawing(question, container);
                    break;
                default:
                    this._renderText(question, container);
            }
        },

        _renderSingleChoice(question, container) {
            const list = document.createElement('div');
            list.className = 'option-list';

            question.options.forEach(function(opt) {
                const item = document.createElement('div');
                item.className = 'option-item';
                item.dataset.optionId = opt.optionId;
                item.innerHTML = '<div class="option-radio"></div><span>' +
                    PlayerRenderer._escape(opt.text) + '</span>';
                item.addEventListener('click', function() {
                    list.querySelectorAll('.option-item').forEach(function(el) {
                        el.classList.remove('selected');
                    });
                    item.classList.add('selected');
                });
                list.appendChild(item);
            });

            container.appendChild(list);
        },

        _renderMultiChoice(question, container) {
            const list = document.createElement('div');
            list.className = 'option-list';

            question.options.forEach(function(opt) {
                const item = document.createElement('div');
                item.className = 'option-item';
                item.dataset.optionId = opt.optionId;
                item.innerHTML = '<div class="option-radio"></div><span>' +
                    PlayerRenderer._escape(opt.text) + '</span>';
                item.addEventListener('click', function() {
                    item.classList.toggle('selected');
                });
                list.appendChild(item);
            });

            container.appendChild(list);
        },

        _renderNumeric(question, container) {
            const input = document.createElement('input');
            input.type = 'text';
            input.className = 'numeric-input';
            input.id = 'answer-input';
            input.placeholder = '?';
            input.inputMode = 'decimal';
            input.autocomplete = 'off';
            container.appendChild(input);
        },

        _renderText(question, container) {
            const input = document.createElement('input');
            input.type = 'text';
            input.className = 'numeric-input';
            input.id = 'answer-input';
            input.placeholder = 'Type your answer...';
            input.style.width = '100%';
            input.style.fontSize = '1.4rem';
            container.appendChild(input);
        },

        _renderDragDrop(question, container) {
            // Placeholder — drag-and-drop requires more complex interaction
            const msg = document.createElement('div');
            msg.className = 'question-prompt';
            msg.textContent = '(Drag and drop activity)';
            container.appendChild(msg);
        },

        _renderDrawing(question, container) {
            // Placeholder — drawing requires canvas
            const msg = document.createElement('div');
            msg.className = 'question-prompt';
            msg.textContent = '(Drawing activity)';
            container.appendChild(msg);
        },

        // Get the student's answer from the current rendered question
        getAnswer(question, container) {
            switch (question.questionType) {
                case 1: { // SINGLE_CHOICE
                    const selected = container.querySelector('.option-item.selected');
                    return selected ? selected.dataset.optionId : '';
                }
                case 2: { // MULTI_CHOICE
                    const selected = container.querySelectorAll('.option-item.selected');
                    return Array.from(selected).map(function(el) { return el.dataset.optionId; }).join(',');
                }
                case 3: // NUMERIC
                case 4: { // TEXT
                    const input = container.querySelector('#answer-input');
                    return input ? input.value.trim() : '';
                }
                default:
                    return '';
            }
        },

        // Show correct/incorrect feedback on options
        showFeedback(question, container, isCorrect) {
            if (question.questionType === 1 || question.questionType === 2) {
                const items = container.querySelectorAll('.option-item');
                items.forEach(function(item) {
                    const opt = question.options.find(function(o) { return o.optionId === item.dataset.optionId; });
                    if (opt && opt.isCorrect) {
                        item.classList.add('correct');
                    } else if (item.classList.contains('selected') && !isCorrect) {
                        item.classList.add('incorrect');
                    }
                });
            }
        },

        _escape(text) {
            var div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }
    };
})();
