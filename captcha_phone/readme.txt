���ŷ��ͷ���

step1��
main.go
�޸ļ�����ַ0.0.0.0-->����ip

step2:
#����docker image
�ڵ�ǰ�ļ��п��նˣ�����
docker build -t message .

step3:
#��������
docker run -p 8666:8666 -d message

END:
#�ӿ�����ӿ��ĵ���