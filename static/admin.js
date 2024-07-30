document.addEventListener("DOMContentLoaded", function() {
    function fetchFlowRates() {
        fetch('/api/flowrates')
            .then(response => response.json())
            .then(data => {
                let html = "<ul>";
                data.forEach(area => {
                    html += `<li>${area.name}: ${area.flow_rate} L/min</li>`;
                });
                html += "</ul>";
                document.getElementById('area-flow-rates').innerHTML = html;
            })
            .catch(error => console.error('Error fetching flow rates:', error));
    }

    function fetchFairDistribution() {
        fetch('/api/fairdistribution')
            .then(response => response.json())
            .then(data => {
                document.getElementById('fair-distribution').textContent = JSON.stringify(data, null, 2);
            })
            .catch(error => console.error('Error fetching fair distribution:', error));
    }

    function fetchLeakDetection() {
        fetch('/api/leakdetection')
            .then(response => response.json())
            .then(data => {
                document.getElementById('leak-detection').textContent = JSON.stringify(data, null, 2);
            })
            .catch(error => console.error('Error fetching leak detection:', error));
    }

    function fetchBlockchain() {
        fetch('/api/blockchain')
            .then(response => response.json())
            .then(data => {
                if (!data) {
                    console.error('Received null data for blockchain');
                    return;
                }
                let html = "<ul>";
                data.forEach(block => {
                    html += `<li>Index: ${block.index}, Timestamp: ${block.timestamp}, Data: ${JSON.stringify(block.data)}, PrevHash: ${block.prevHash}, Hash: ${block.hash}</li>`;
                });
                html += "</ul>";
                document.getElementById('blockchain-data').innerHTML = html;
            })
            .catch(error => console.error('Error fetching blockchain data:', error));
    }

    // Fetch data initially and then periodically
    fetchFlowRates();
    fetchFairDistribution();
    fetchLeakDetection();
    fetchBlockchain();
    setInterval(() => {
        fetchFlowRates();
        fetchFairDistribution();
        fetchLeakDetection();
        fetchBlockchain();
    }, 5000); // Update every 5 seconds
});
