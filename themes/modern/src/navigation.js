import './spatial_navigation'

// Define navigable elements (anchors and elements with "focusable" class).
const sections = { };
function addSection(id, parent) {
  if(SpatialNavigation.has(id)) return false;
  sections[id] = parent;

  SpatialNavigation.add(id, {
    selector: '.focusable-' + id + (id == 'main' ? ', a' : ''),
    rememberSource: true
  });

  return true;
}

// Initialize
SpatialNavigation.init();

addSection('main');

// Play the sound when an anchor or focusable is clicked.
// This may be changed in the future to just anchors, with
// other elements playing their own custom sounds.
document.addEventListener('click', e => {
    if(!e.target.tagName) return;
    
    if(e.target.tagName.toLowerCase() !== 'a'
        && e.target.classList.toString().indexOf('focusable') === -1)
        return;

    new Audio('/sounds/test4.mp3').play();
}, false);

let lastPop = null;
function bindListener(el) {
    ['mouseover', 'sn:focused'].forEach(event => {
        el.addEventListener(event, e => {
            // For "document" we want to use the element that was dragged over.
            // For others, we want the bound element, because children will bubble
            // the event up to it.
            let target = el != document ? el : e.target;

            if(!target.tagName) return;

            if(target.tagName.toLowerCase() !== 'a'
                && target.classList.toString().indexOf('focusable') === -1)
                return;

            // Set the last pop variable so we don't play a sound again when a
            // child of the element is interacted with
            if(lastPop == target) return;
            lastPop = target;
        
            new Audio('/sounds/test6.mp3').play();

            // If we don't do this, then the mouse will be impossible to use.
            if(event == 'sn:focused')
                e.target.scrollIntoView({ behavior: 'smooth', block: 'center', inline: 'center' });
        }, true);
    });
}

bindListener(document);

export default {
    install(Vue) {
        Vue.directive('focus-section', {
            bind(el, binding) {
                let sectionId = binding.value || 'main';

                if(addSection(sectionId, el))
                console.log('Added section:', sectionId);
            }
        });

        Vue.directive('focus', {
            inserted(el, binding, vnode) {
                if(el.classList.toString().indexOf('focusable-') === -1) {
                    let sectionId = Object.keys(binding.modifiers)[0] || null;

                    if(sectionId == null) {
                        let newSection = 'main';

                        for(let [id, parent] of Object.entries(sections)) {
                            if(!parent) continue;

                            if(parent.contains(el)) {
                                newSection = id;
                                break;
                            }
                        }

                        sectionId = newSection;
                    }

                    el.classList.add('focusable');
                    el.classList.add('focusable-' + sectionId);

                    addSection(sectionId);

                    SpatialNavigation.makeFocusable(sectionId);
                }

                bindListener(el);

                el.addEventListener('sn:enter-down', () => {
                    if(el.tagName.toLowerCase() == 'video') {
                        // We should play/pause videos instead of emulating a click
                        if(!!(el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2)) {
                            el.pause();
                        }else{
                            el.play();
                        }
                    }else{
                        if(vnode.componentInstance) {
                            vnode.componentInstance.$emit('click', { });
                        }else{
                            vnode.elm.dispatchEvent(new CustomEvent('click', {}));
                        }
                    }
                });
            }
        });
    }
}