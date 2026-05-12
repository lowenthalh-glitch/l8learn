/*
 * L8Learn Student Player — Renderer
 * Renders lesson steps and questions into the activity container
 */
(function() {
    'use strict';

    window.PlayerRenderer = {

        // Render a full lesson step into the container
        renderStep(step, container) {
            container.innerHTML = '';
            var type = (step.stepType || '').toLowerCase();

            if (type === 'physical' || type === 'hands-on') {
                this._renderPhysicalStep(step, container);
            } else if (type === 'worksheet') {
                this._renderWorksheetStep(step, container);
            } else {
                // "screen" or unknown — render as question-based step
                this._renderScreenStepHeader(step, container);
            }
        },

        // Physical step: show instructions + materials + "Mark Complete" button
        _renderPhysicalStep(step, container) {
            var card = document.createElement('div');
            card.className = 'step-physical';

            var icon = document.createElement('div');
            icon.className = 'step-type-badge physical';
            icon.textContent = 'Hands-On Activity';
            card.appendChild(icon);

            var title = document.createElement('h3');
            title.className = 'step-title';
            title.textContent = step.title || 'Activity';
            card.appendChild(title);

            var instructions = document.createElement('div');
            instructions.className = 'step-instructions';
            instructions.textContent = step.instructions || '';
            card.appendChild(instructions);

            if (step.materialsInstructions) {
                var matLabel = document.createElement('div');
                matLabel.className = 'step-materials-label';
                matLabel.textContent = 'What you need:';
                card.appendChild(matLabel);

                var matText = document.createElement('div');
                matText.className = 'step-materials';
                matText.textContent = step.materialsInstructions;
                card.appendChild(matText);
            }

            if (step.parentRole && step.parentRole !== 'none') {
                var parentNote = document.createElement('div');
                parentNote.className = 'step-parent-note';
                parentNote.textContent = 'Parent: ' + step.parentRole;
                card.appendChild(parentNote);
            }

            if (step.durationMinutes) {
                var dur = document.createElement('div');
                dur.className = 'step-duration';
                dur.textContent = 'About ' + step.durationMinutes + ' minutes';
                card.appendChild(dur);
            }

            container.appendChild(card);
        },

        // Worksheet step: show instructions and download link
        _renderWorksheetStep(step, container) {
            var card = document.createElement('div');
            card.className = 'step-worksheet';

            var icon = document.createElement('div');
            icon.className = 'step-type-badge worksheet';
            icon.textContent = 'Worksheet';
            card.appendChild(icon);

            var title = document.createElement('h3');
            title.className = 'step-title';
            title.textContent = step.title || 'Practice Worksheet';
            card.appendChild(title);

            var instructions = document.createElement('div');
            instructions.className = 'step-instructions';
            instructions.textContent = step.instructions || 'Complete the worksheet below.';
            card.appendChild(instructions);

            if (step.durationMinutes) {
                var dur = document.createElement('div');
                dur.className = 'step-duration';
                dur.textContent = 'About ' + step.durationMinutes + ' minutes';
                card.appendChild(dur);
            }

            container.appendChild(card);
        },

        // Screen step header (questions are rendered individually after)
        _renderScreenStepHeader(step, container) {
            if (step.title) {
                var title = document.createElement('div');
                title.className = 'step-type-badge screen';
                title.textContent = step.title || 'Practice';
                container.appendChild(title);
            }
            if (step.instructions) {
                var instr = document.createElement('div');
                instr.className = 'step-instructions';
                instr.textContent = step.instructions;
                container.appendChild(instr);
            }
        },

        // Render a GeneratedQuestion into the container
        renderQuestion(question, container) {
            // Remove any previous question content but keep step header
            var existing = container.querySelectorAll('.question-prompt, .option-list, .numeric-input, .question-media, .question-hint');
            existing.forEach(function(el) { el.remove(); });

            var prompt = document.createElement('div');
            prompt.className = 'question-prompt';
            // GeneratedQuestion uses "prompt", legacy Activity uses "promptText"
            prompt.textContent = question.prompt || question.promptText || '';
            container.appendChild(prompt);

            if (question.promptMediaUrl || question.promptImageDesc) {
                var media = document.createElement('div');
                media.className = 'question-media';
                if (question.promptMediaUrl) {
                    var img = document.createElement('img');
                    img.src = question.promptMediaUrl;
                    img.alt = question.promptImageDesc || 'Question image';
                    media.appendChild(img);
                } else if (question.promptImageDesc) {
                    media.textContent = question.promptImageDesc;
                    media.style.fontStyle = 'italic';
                    media.style.color = 'var(--player-text-light)';
                }
                container.appendChild(media);
            }

            switch (question.questionType) {
                case 1: this._renderSingleChoice(question, container); break;
                case 2: this._renderMultiChoice(question, container); break;
                case 3: this._renderNumeric(question, container); break;
                case 4: this._renderText(question, container); break;
                case 5: this._renderDragDrop(question, container); break;
                case 6: this._renderDrawing(question, container); break;
                default: this._renderText(question, container);
            }
        },

        _renderSingleChoice(question, container) {
            var list = document.createElement('div');
            list.className = 'option-list';
            (question.options || []).forEach(function(opt) {
                var item = document.createElement('div');
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
            var list = document.createElement('div');
            list.className = 'option-list';
            (question.options || []).forEach(function(opt) {
                var item = document.createElement('div');
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
            var input = document.createElement('input');
            input.type = 'text';
            input.className = 'numeric-input';
            input.id = 'answer-input';
            input.placeholder = '?';
            input.inputMode = 'decimal';
            input.autocomplete = 'off';
            container.appendChild(input);
        },

        _renderText(question, container) {
            var input = document.createElement('input');
            input.type = 'text';
            input.className = 'numeric-input';
            input.id = 'answer-input';
            input.placeholder = 'Type your answer...';
            input.style.width = '100%';
            input.style.fontSize = '1.4rem';
            container.appendChild(input);
        },

        _renderDragDrop(question, container) {
            var msg = document.createElement('div');
            msg.className = 'question-prompt';
            msg.textContent = '(Drag and drop activity)';
            container.appendChild(msg);
        },

        _renderDrawing(question, container) {
            var msg = document.createElement('div');
            msg.className = 'question-prompt';
            msg.textContent = '(Drawing activity)';
            container.appendChild(msg);
        },

        // Get the student's answer from the current rendered question
        getAnswer(question, container) {
            switch (question.questionType) {
                case 1: {
                    var selected = container.querySelector('.option-item.selected');
                    return selected ? selected.dataset.optionId : '';
                }
                case 2: {
                    var items = container.querySelectorAll('.option-item.selected');
                    return Array.from(items).map(function(el) { return el.dataset.optionId; }).join(',');
                }
                case 3:
                case 4: {
                    var input = container.querySelector('#answer-input');
                    return input ? input.value.trim() : '';
                }
                default:
                    return '';
            }
        },

        // Show correct/incorrect feedback on options
        showFeedback(question, container, isCorrect) {
            if (question.questionType === 1 || question.questionType === 2) {
                var items = container.querySelectorAll('.option-item');
                items.forEach(function(item) {
                    var opt = (question.options || []).find(function(o) {
                        return o.optionId === item.dataset.optionId;
                    });
                    if (opt && opt.isCorrect) {
                        item.classList.add('correct');
                    } else if (item.classList.contains('selected') && !isCorrect) {
                        item.classList.add('incorrect');
                    }
                });
            }

            // Show explanation if available
            if (question.explanation) {
                var expl = document.createElement('div');
                expl.className = 'question-hint';
                expl.style.cssText = 'background:#e0f2fe;padding:12px;border-radius:8px;margin-top:12px;font-size:1rem;';
                expl.textContent = question.explanation;
                container.appendChild(expl);
            }
        },

        // Render a hint for a GeneratedQuestion
        showHint(question, hintIndex, container) {
            // GeneratedQuestion uses "hints", legacy uses "hintTexts"
            var hints = question.hints || question.hintTexts || [];
            if (hintIndex < hints.length) {
                var hintEl = document.createElement('div');
                hintEl.className = 'question-hint';
                hintEl.style.cssText = 'background:#fef3c7;padding:12px;border-radius:8px;margin-top:12px;font-size:1rem;';
                hintEl.textContent = hints[hintIndex];
                container.appendChild(hintEl);
                return true;
            }
            return false;
        },

        // Check if question has hints available
        hasHints(question) {
            var hints = question.hints || question.hintTexts || [];
            return hints.length > 0;
        },

        _escape(text) {
            var div = document.createElement('div');
            div.textContent = text || '';
            return div.innerHTML;
        }
    };
})();
