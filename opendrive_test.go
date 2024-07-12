package xodr

const ExampleXodr = `<OpenDRIVE>
<header name="mt/set_7/frenetic/odr/test-15">
<sdc_test_info test_id="mt/set_7/frenetic/odr/test-15" test_outcome="FAIL" predicted_test_outcome="null" test_duration="8.018114805221558" is_valid="True" />
<run_config rf="1.5" oob="0.5" max_speed="120" obstacles="False" bump_dist="20" delineator_dist="None" tree_dist="None" field_of_view="120" />
</header>
<road id="15" length="1.2970567222335126e+02">
<link />
<planView>
<geometry s="0.0000000000000000e+00" x="6.4836535755501515e+01" y="1.0000000000000004e+01" hdg="0.0000000000000000e+00" length="1.2955170009307626e+00">
<paramPoly3 pRange="normalized" aU="0.0000000000000000e+00" bU="6.8078264457398063e-02" cU="0.0000000000000000e+00" dU="0.0000000000000000e+00" aV="0.0000000000000000e+00" bV="1.2904538720509038e+00" cV="0.0000000000000000e+00" dV="0.0000000000000000e+00" />
</geometry>
<geometry s="1.2955170009307626e+00" x="6.4904764449067358e+01" y="1.1293719113669944e+01" hdg="0.0000000000000000e+00" length="6.5096045624843246e-01">
<paramPoly3 pRange="normalized" aU="-1.4210854715202004e-14" bU="2.9773957985937190e-02" cU="-2.1792302961051876e-03" dU="-2.8607423949722488e-05" aV="0.0000000000000000e+00" bV="6.4652087482110454e-01" cV="6.0067343101834691e-04" dV="-2.1468083597807498e-05" />
</geometry>
<geometry s="1.9464774571791952e+00" x="6.4932458439205618e+01" y="1.1944090206592584e+01" hdg="0.0000000000000000e+00" length="6.5186386720529588e-01">
<paramPoly3 pRange="normalized" aU="-1.4210854715202004e-14" bU="2.5306793613889111e-02" cU="-2.2654860137737848e-03" dU="-2.8607424016039811e-05" aV="-1.7763568394002505e-15" bV="6.4766323264995629e-01" cV="5.3594390623057267e-04" dV="-2.1468083612511773e-05" />
</geometry>
<geometry s="1.2838922362470572e+02" x="1.1971088945308670e+01" y="1.0639669813112674e+02" hdg="0.0000000000000000e+00" length="6.5717042837793993e-01">
<paramPoly3 pRange="normalized" aU="-1.7763568394002505e-15" bU="-6.5234979737857079e-01" cU="-4.2460823780969087e-04" dU="-1.0772727705961790e-05" aV="1.4210854715202004e-14" bV="2.9561073219503992e-02" cV="7.7803677363831010e-03" dV="1.9407884956379462e-04" />
</geometry>
<geometry s="1.2904639405308367e+02" x="1.1315004607166498e+01" y="1.0643446469277846e+02" hdg="0.0000000000000000e+00" length="6.5927817026758917e-01">
<paramPoly3 pRange="normalized" aU="1.7763568394002505e-15" bU="-6.5323594828004639e-01" cU="-4.5708964406027586e-04" dU="-1.0772727701126157e-05" aV="0.0000000000000000e+00" bV="4.5788530842638980e-02" cV="8.3655448735682567e-03" dV="1.9407884951800429e-04" />
</geometry>
</planView>
<lanes>
<laneSection s="0.0000000000000000e+00">
<left>
<lane id="1" type="driving" level="False">
<link />
<width a="3.0000000000000000e+00" />
<roadMark type="solid" width="2.0000000000000001e-01" />
</lane>
</left>
<center>
<lane id="0" type="none" level="False">
<link />
<width a="3.0000000000000000e+00" />
<roadMark type="solid" width="2.0000000000000001e-01" />
</lane>
</center>
<right>
<lane id="-1" type="driving" level="False">
<link />
<width a="3.0000000000000000e+00" />
<roadMark type="solid" width="2.0000000000000001e-01" />
</lane>
</right>
</laneSection>
</lanes>
</road>
</OpenDRIVE>`
