const searchInput = document.getElementById('searchInput');
const suggestionsBox = document.getElementById('suggestionsBox');

searchInput.addEventListener('input', (e) => {
    const query = e.target.value.toLowerCase();
    suggestionsBox.innerHTML = ''; // Vider les suggestions précédentes

    if (query.length > 0) {
        // Filtrer le tableau rempli par Go
        const filteredArtists = artistsData.filter(artist => {
            return artist.name.toLowerCase().includes(query);
            // Vous pouvez ajouter ici la recherche sur les membres ou les dates
        });

        // Afficher les suggestions
        filteredArtists.forEach(artist => {
            const div = document.createElement('div');
            div.textContent = artist.name;
            div.classList.add('suggestion-item');

            // Gérer le clic sur une suggestion (rediriger ou afficher)
            div.addEventListener('click', () => {
                window.location.href = `/artist/${artist.id}`; // Exemple de redirection
            });

            suggestionsBox.appendChild(div);
        });
    }
});