# Algoritma AI

## Tipe-Tipe AI
#### Berdasarkan level kecerdasan :
* Narrwo/Weak AI
* General/Strong AI
* Super AI

#### Berdasarkan Fungsionalitas:
* Reactive Machine
* Limited Memory
* Theory of Mind
* Self - Awareness

Pada saat ini, tipe A.I. yang umum digunakan adalah narrow A.I. atau reactive machine dan limited memory. Tipe lain masih sebatas teori atau dalam tahap pengembangan

## Narrow/Weak AI
* AI untuk melaksanakan tugas yang spesifik
* Kemampuannya sangat dioptimaliasi untuk sebuah tugas dan dapat melampaui kemampuan manusia
Contoh :
* Facial Recognition
* Natural Language Processing (ChatGPT)
* Autonomous Vehicle
* voice Assistants (Siri, Alexa)

## Reactive Machine & Limited Memory
Saat ini, reactive machine dan limited memory AI sudah diciptakan
* Reactive machine : tipe AI yang melakukan reaksi terhadap input tanpa pengetahuan masa lampau dan hanya spesifik melakukan tugas tertentu. Contoh : Machine learning, recomnendation engine, Deep blue (Mesin catur IBM)
* Limited memory : pengembangan dari reactive machine, sudah dilengkapi dengan kemampuan mengevaluasi situasi masa lampau. Contoh: deep learning , reinforcement learning, self-drivinf cars

## Definisi Algoritma
* Sebuah set untruksi yang diberikan kepada komputer untuk memecahkan masalah atau mencapai sebuah tujuan, contoh : menghitung luas segitiga
* Dalam konteks AI, algoritma digunakan agar komputer mampu melakukan kalkulasi dan pemrosesan data, melatih model machine learning dan deep learning, dan membuat prediksi berdasarkan data yang diberikan

## Machine Learning
* Subset dari AI yang fokus pada pengembangan algoritma dan model
* Memiliki kemampuan untuk belajar, dan mengambil keputusan atau prediksi tanpa perlu diprogram secara 
* Kualitas dari model machine learning tergantung dari bagus/tidaknya data yang diberikan

## Kategori Machine Learning
## Supervised Learning
* Model machine learning di training pada data berlabel, dimana nilai input - output sudah diketahui
* Algoritma machine learning mempelajari cara me-mapping sebuah function dari input untuk menghasilkan output, Y = f(X)
* Ada dua tipe supervised learning : Regression dan Classification

### Regression
Output berupa variabel kontinu (angka)
* Linear regression
* Polynimial regression
* Logistic regression

### Classification
Output berupa kategori
* Binary classification
* Multi-class classification
* Multi-label classification

## Contoh Algoritma Supervised Learning
### Regression : 
* Linear Regression
* Polynimial Regression
* Support Vector Regression
* Decision Tree Regression
* Random Forest Regression

### Classification :
* Logistic Regression
* Decision Tree
* Support Vector Machines (SVM)
* K-Nearest Neighbors (K-NN)
* Random Forest
* Naive Bayes

## Unsupervised learning
* Model machine learning di training pada data tidak berlabel
* Mencari pola, mengelompokan data berdasarkan kesamaan atribut pada dataset yang diberikan dan membuat keputusan berdasarkan apa yang ditemukan
* Ada 3 tipe Unsupervised learning : Clustering, Dimensionality Reduction dan Anomaly Detection

### Clustering
* K-means Clustering
* Hierarchical Clustering
* Density-based Clustering

### Dimensionality Reduction
* Principal Component Analysis (PCA)
* Indipendent Component Analysis (ICA)
* Non-negative Matrix Factorization (NMF)

### Anomaly Detection
* Statistical-based
* Distance-based
* Density-based

## Contoh Algoritma Unsupervised Learning
### Clustering
* K-means
* Density-based (DBSCAN)
* Gaussian Mixture Model
* BIRCH
* Afinity Propagation
* Mean-shift
* OPTICS
* Agglomerative Hierarchy

### Dimensionality Reduction
* Principal Component Analysis (PCA)
* Singular Value Decomposition (SVD)
* Linier Discrimination Analysis (LDA)
* Non-negative Matrix Factorization (NMF)
* T-distributed Stochastic Neighbor Embedding (t-sne)
* Factor Analysis
* Independent Component Analysis (ICA)

### Anomaly Detection
* Isolation Forest
* One-class SVM
* Local Outlier Factor (LOF)

### Association Rule Learning
* Apriori Algorithm
* Eclat Algorithm
* Frequent Pattern(FP)-Growth Algorithm

## Semi-Supervised Learning
* Perpaduan antara supervised learning dengan semi-supervised learning
* Mengombinasikan data berlabel dan tidak berlabel untuk meningkatkan akurasi model machine learning
* Digunakan saat sejumlah besar data yang tersedia tidak berlabel, namun diperlukan waktu yang lama atau biaya yang mahal untuk memberi label
* Self learning, co-training, graph-based method, multi-view learning

## Contoh Algoritma Semi-Supervised Learning
* Support Vector Machines (SVM)
* K-nearest Neightbors
* Decision Trees
* Label Propagation
* Label Spreading
* Multi-view Clustering
* Multi-view Dimensionality Reduction
* Multi-view Layering

## Reinforcement Learning
* Sebuah proses iteratif dimana agent berinteraksi dengan environment dan belajar dari feedback yang didapat
* Belajar dengan cara memaksimalkan reward dan meminimalkan punishment dalam rangka mencapai goal yang diberikan
* Model-based, Value-based, Policy-based

## Contoh Algoritma Reinforcement Learning
* Dyna
* Policy Gradient
* Direct Utility Estimation
* Q-learning
* Temporal-Different (TD) Learning
* Adaptive Dynamic Programming (ADP)
* SARSA

## Deep Learning
* Subset dari machine learning
* Menggunakan artificial neural network untuk meniru proses pembelajaran manusia
* Dibandingkan machine learning, dapat membangun korelasi antara data dengan lebih kompleks, namun membutuhkan dataset yang relatif lebih besar
* Idealnya menggunakan GPU (graphic processing unit) dalam melatih model deep learning

## Machine Learning vs Deep Learning
Machine Learning
* Input > Feature Extraction > Classification > Output
Deep Learning
* Input > Feature Extraction + Classification > Ouput

Deep Learning tidak membutuhkan feature extraction secara manual

## Contoh Algoritma Deep Learning
* Convolutional Neural Networks (CNNs)
* Recurrent Neural Networks (RNNs)
* Long Short-Term Memory Networks (LSTMs)
* Generative Adversarial Networks (GANs)
* Radial Basis Function Networks (RBFNs)
* Multilayer Perceptrons (MLPs)
* Self Organizing Maps (SOMs)
* Deep Belief Networks(DBNs)
* Restricted Boltzmann Machines (RBMs)
* Autoencoders

## Frameworks & Libraries Untuk AI
### Pyton
* Pandas
* TensorFlow
* Keras
* NumPy
* Scikit-learn
* PyTorch

### Java
* Deep Java Library (DJL)
* Deeplearning4j
* Apache OpenNLP
* Java Machine Learning Library (Java-ML)
* RapidMiner
* Weka
* Encog Machine Learning Framework

### Go
* golearn
* gorgonia
* eaopt
* gonum
* gomind

### Flutter
* flutter_openai
* egde_detection
* tflite_flutter
* google_ml_kit