import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  main: {
    async getList(): Promise<string[]> {
      return (await Axios.get(`${API_URL}/main/list`)).data.response;
    },
    async getConfig(): Promise<any> {
      return (await Axios.get(`${API_URL}/main/config`)).data.response;
    },
    async updateConfig({ SPACES_KEY, SPACES_SECRET, SPACES_ENDPOINT, SPACES_BUCKET }: any) {
      return (
        await Axios.post(`${API_URL}/main/config`, {
          SPACES_KEY,
          SPACES_SECRET,
          SPACES_ENDPOINT,
          SPACES_BUCKET,
        })
      ).data.response;
    },
    async backup(path: string) {
      return (
        await Axios.post(`${API_URL}/main/backup`, {
          path,
        })
      ).data.response;
    },
    async getBackupList(path: string): Promise<string[]> {
      return (await Axios.get(`${API_URL}/main/backupList?path=${path}`)).data.response;
    },
  },
};
