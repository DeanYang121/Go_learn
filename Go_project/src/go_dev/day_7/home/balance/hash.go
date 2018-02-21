package balance
import(
	"hash/crc32"
	"fmt"
	"math/rand"
)

type HashBalance struct{

}

func init(){
	RegisterBalancer("hash",&HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error){
	
	var defKey string = fmt.Sprintf("%d",rand.Int())

	if len(key) > 0{
		defKey = key[0]
		//err = fmt.Errorf("hash balance must pass hash key")
		return
	}

	lens := len(insts)

	if lens == 0{
		err = fmt.Errorf("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey),crcTable)
	index := int(hashVal) % lens

	inst = insts[index]
	return 
}