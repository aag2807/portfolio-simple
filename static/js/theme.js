(function() {
    var toggle = document.getElementById('theme-toggle');
    var html = document.documentElement;
    
    var stored = localStorage.getItem('theme');
    if (stored === 'dark' || (!stored && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        html.classList.add('dark');
    }
    
    toggle.addEventListener('click', function() {
        html.classList.toggle('dark');
        localStorage.setItem('theme', html.classList.contains('dark') ? 'dark' : 'light');
    });
})();

