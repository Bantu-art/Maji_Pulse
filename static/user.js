function fetchUserData() {
    fetch('/user-data')
        .then(response => response.json())
        .then(data => {
            console.log("Fetched data:", data); // Log the data to debug

            // Check if data is an array and process the first item
            if (Array.isArray(data) && data.length > 0) {
                const item = data[0]; // Get the first item in the array
                if (item.flow_rate !== undefined && item.cost !== undefined) {
                    document.querySelector('#flow-rate p').textContent = `Flow Rate: ${item.flow_rate.toFixed(2)} L/s`;
                    // Update other parts of the page as needed
                } else {
                    console.error('Invalid data item:', item);
                }
            } else {
                console.error('Invalid data received:', data);
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

// Fetch data every second
setInterval(fetchUserData, 1000);
fetchUserData();  // Initial fetch