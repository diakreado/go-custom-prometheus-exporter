import axios from 'axios';

const prmetheusAddr = 'http://localhost192.168.0.47:9090';

(async () => {
	const response = await axios.get(`${prmetheusAddr}/api/v1/query?query=my_app_ingress_application`);
	console.log(response.data.data.result);
})()
