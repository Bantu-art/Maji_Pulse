function fetchAdminData() {
    fetch('/admin-data')
        .then(response => response.json())
        .then(data => {
            console.log("Fetched data:", data); // Log the data to debug

            if (Array.isArray(data)) {
                data.forEach(item => {
                    if (item.area && item.flow_rate !== undefined && item.leakage !== undefined && item.fair_percentage !== undefined) {
                        const areaID = `flowrate-${item.area.replace(/ /g, '').toLowerCase()}`; // Dynamically create the ID
                        const leakageID = `leakage-${item.area.replace(/ /g, '').toLowerCase()}`; // Dynamically create the ID
                        const distributionID = `distribution-${item.area.replace(/ /g, '').toLowerCase()}`; // Dynamically create the ID

                        // Update flow rate section
                        const flowRateElement = document.getElementById(areaID);
                        if (flowRateElement) {
                            flowRateElement.querySelector('p').textContent = `Flow Rate: ${item.flow_rate.toFixed(2)} L/s`;
                        }

                        // Update leakage section
                        const leakageElement = document.getElementById(leakageID);
                        if (leakageElement) {
                            leakageElement.querySelector('p').textContent = `Leakage: ${item.leakage ? 'Yes' : 'No'}`;
                            leakageElement.querySelector('p').className = item.leakage ? 'leakage-alert' : '';
                        }

                        // Update distribution section
                        const distributionElement = document.getElementById(distributionID);
                        if (distributionElement) {
                            distributionElement.querySelector('p').textContent = `Fair Percentage: ${item.fair_percentage.toFixed(2)}%`;
                        }
                    } else {
                        console.error('Invalid data item:', item);
                    }
                });
            } else {
                console.error('Invalid data received:', data);
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

// Fetch data every second
setInterval(fetchAdminData, 1000);
fetchAdminData();  // Initial fetch