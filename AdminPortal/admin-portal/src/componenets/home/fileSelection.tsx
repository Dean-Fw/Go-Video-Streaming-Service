import { useState } from 'react';
import PostNewResource from '../../features/uploadServiceApi/postNewResource';
import PostNewReosourceHeadersModel from '../../features/uploadServiceApi/requestModels/postNewResourceHeadersModel';
import styles from './styles/fileSelection.module.css'

const FileSelection = () => {
	const [videoFile, setVideoFile] = useState<File>();
	const [showError, setShowError] = useState<boolean>(false);
	const [errorMessage, setErrorMesasge] = useState<string>("");

	const onFileChange = (event: React.FormEvent) => {
		const files = (event.target as HTMLInputElement).files

		if (files && files.length > 0) {
			setVideoFile(files[0])
		}
	}

	const onButtonPress = () => {
		if (videoFile == null) {
			setErrorMesasge("Please select a file to upload")
			setShowError(true)
		}
		else {
			setErrorMesasge("")
			setShowError(false)
			console.log(videoFile?.name)
			CreateNewResource()
		}
	}

	const CreateNewResource = () => {
		const postNewResourceHeaders: PostNewReosourceHeadersModel = {
			FileName: videoFile!.name,
			UploadLength: videoFile!.size,
			ContentLength: 10 ^ 6
		}
		PostNewResource(postNewResourceHeaders)
	}

	return (
		<div className={styles.fileSelection}>
			{showError ? <p>{errorMessage}</p> : <></>}
			<input type="file" accept="video/mp4" onChange={onFileChange} />
			<button onClick={onButtonPress}> UPLOAD! </button>
		</div>
	)
}

export default FileSelection
