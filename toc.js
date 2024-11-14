// Populate the sidebar
//
// This is a script, and not included directly in the page, to control the total size of the book.
// The TOC contains an entry for each page, so if each page includes a copy of the TOC,
// the total size of the page becomes O(n**2).
class MDBookSidebarScrollbox extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.innerHTML = '<ol class="chapter"><li class="chapter-item "><a href="00-overview.html"><strong aria-hidden="true">1.</strong> Overview</a></li><li class="chapter-item "><a href="01-declaring-models.html"><strong aria-hidden="true">2.</strong> Declaring Models</a></li><li class="chapter-item "><a href="02-connecting-to-database.html"><strong aria-hidden="true">3.</strong> Connecting to a Database</a></li><li class="chapter-item "><a href="03-crud/index.html"><strong aria-hidden="true">4.</strong> CRUD Interface</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="03-crud/01-create.html"><strong aria-hidden="true">4.1.</strong> Create</a></li><li class="chapter-item "><a href="03-crud/02-query.html"><strong aria-hidden="true">4.2.</strong> Query</a></li><li class="chapter-item "><a href="03-crud/03-advanced-query.html"><strong aria-hidden="true">4.3.</strong> Advanced Query</a></li><li class="chapter-item "><a href="03-crud/04-update.html"><strong aria-hidden="true">4.4.</strong> Update</a></li><li class="chapter-item "><a href="03-crud/05-delete.html"><strong aria-hidden="true">4.5.</strong> Delete</a></li><li class="chapter-item "><a href="03-crud/06-raw-sql.html"><strong aria-hidden="true">4.6.</strong> Raw SQL &amp; SQL Builder</a></li></ol></li><li class="chapter-item "><a href="04-original-gorm-db.html"><strong aria-hidden="true">5.</strong> Original gorm db</a></li><li class="chapter-item "><a href="05-transaction.html"><strong aria-hidden="true">6.</strong> Transaction</a></li><li class="chapter-item "><a href="06-associations.html"><strong aria-hidden="true">7.</strong> Associations</a></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString();
        if (current_page.endsWith("/")) {
            current_page += "index.html";
        }
        var links = Array.prototype.slice.call(this.querySelectorAll("a"));
        var l = links.length;
        for (var i = 0; i < l; ++i) {
            var link = links[i];
            var href = link.getAttribute("href");
            if (href && !href.startsWith("#") && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The "index" page is supposed to alias the first chapter in the book.
            if (link.href === current_page || (i === 0 && path_to_root === "" && current_page.endsWith("/index.html"))) {
                link.classList.add("active");
                var parent = link.parentElement;
                if (parent && parent.classList.contains("chapter-item")) {
                    parent.classList.add("expanded");
                }
                while (parent) {
                    if (parent.tagName === "LI" && parent.previousElementSibling) {
                        if (parent.previousElementSibling.classList.contains("chapter-item")) {
                            parent.previousElementSibling.classList.add("expanded");
                        }
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', function(e) {
            if (e.target.tagName === 'A') {
                sessionStorage.setItem('sidebar-scroll', this.scrollTop);
            }
        }, { passive: true });
        var sidebarScrollTop = sessionStorage.getItem('sidebar-scroll');
        sessionStorage.removeItem('sidebar-scroll');
        if (sidebarScrollTop) {
            // preserve sidebar scroll position when navigating via links within sidebar
            this.scrollTop = sidebarScrollTop;
        } else {
            // scroll sidebar to current active section when navigating via "next/previous chapter" buttons
            var activeSection = document.querySelector('#sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        var sidebarAnchorToggles = document.querySelectorAll('#sidebar a.toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(function (el) {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define("mdbook-sidebar-scrollbox", MDBookSidebarScrollbox);
