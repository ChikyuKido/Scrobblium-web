import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: '/api/v1',
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('jwt')}`
    }
})


axiosInstance.interceptors.response.use(
    response => response,
    async error => {
        return Promise.reject(error)
    }
)

export default axiosInstance
