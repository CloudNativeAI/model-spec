/*
 *     Copyright 2024 The CNAI Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

const (
	// AnnotationConfig is the annotation key for the layer is a configuration file (boolean), such as `true` or `false`.
	AnnotationConfig = "org.cnai.model.config"

	// AnnotationModel is the annotation key for the layer is a model file (boolean), such as `true` or `false`.
	AnnotationModel = "org.cnai.model.model"

	// AnnotationFilepath is the annotation key for the file path of the layer.
	AnnotationFilepath = "org.cnai.model.filepath"
)
