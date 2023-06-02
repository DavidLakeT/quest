const API_BASE_URL = 'http://127.0.0.1:3001/apis/document';

export const uploadFile = async (uploadData) => {
    try {
      console.log(JSON.stringify(uploadData))
      const response = await fetch(`${API_BASE_URL}/uploadDocument`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(uploadData)
      });
  
      if (!response.ok) {
        throw new Error('Error uploading document');
      }
  
      const upload = await response.json();
      return upload;
    } catch (error) {
      throw new Error('Error uploading document');
    }
};
