const cache = 'side-challenge-3000-v5-score';
const files = ['./', './index.html', './manifest.webmanifest', './icon-192.png', './icon-512.png'];

self.addEventListener('install', e => {
    e.waitUntil(caches.open(cache).then(c => c.addAll(files)));
    self.skipWaiting();
});

self.addEventListener('activate', e => {
    e.waitUntil(caches.keys().then(a => Promise.all(a.filter(k => k !== cache).map(k => caches.delete(k)))));
    self.clients.claim();
});

self.addEventListener('fetch', e => {
    if (e.request.method !== 'GET') return;
    e.respondWith(caches.match(e.request).then(c => c || fetch(e.request)));
});
